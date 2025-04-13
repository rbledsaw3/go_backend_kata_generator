package main

import (
    "crypto/rand"
    "encoding/binary"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

type Table struct {
    Name            string      `json:"name"`
    Columns         []string    `json:"columns"`
    Dependencies    []string    `json:"dependencies"`
}

type Kata struct {
    ProjectName string              `json:"project_name"`
    Tables      map[string][]string `json:"tables"`
}

func main() {
    adjectives, err := loadWords("adjectives.txt")
    if err != nil {
        log.Fatalf("Error loading adjectives: %v", err)
        os.Exit(1)
    }

    nouns, err := loadWords("nouns.txt")
    if err != nil {
        log.Fatalf("Error loading nouns: %v", err)
        os.Exit(1)
    }

    tables, err := loadTablesFromJSON("schema_templates.json")
    if err != nil {
        log.Fatalf("Error loading tables from JSON: %v", err)
        os.Exit(1)
    }

    projectName := fmt.Sprintf("%s-%s", randomChoice(adjectives), randomChoice(nouns))

    tableMap := make(map[string]Table)
    for _, t := range tables {
        tableMap[t.Name] = t
    }

    selectedTables := make(map[string]Table)

    addTableWithDependencies(tableMap, selectedTables, "users")

    var possibleTables []string
    for _, t := range tables {
        if t.Name != "users" {
            possibleTables = append(possibleTables, t.Name)
        }
    }

    sShuffle(possibleTables)

    for _, tableName := range possibleTables {
        if len(selectedTables) >= 5 {
            break
        }
        addTableWithDependencies(tableMap, selectedTables, tableName)
    }

    finalTables := make(map[string][]string)
    for name, table := range selectedTables {
        finalTables[name] = table.Columns
    }

    kata := Kata{
        ProjectName: projectName,
        Tables:      finalTables,
    }

    kataDir := filepath.Join("..", "katas", projectName)
    err = os.MkdirAll(kataDir, os.ModePerm)
    if err != nil {
        log.Fatalf("Error creating kata directory: %v", err)
        os.Exit(1)
    }

    jsonPath := filepath.Join(kataDir, fmt.Sprintf("%s.json", projectName))
    out, err := json.MarshalIndent(kata, "", "  ")
    if err != nil {
        log.Fatalf("Error marshaling kata to JSON: %v", err)
        os.Exit(1)
    }

    err = os.WriteFile(jsonPath, out, 0644)
    if err != nil {
        log.Fatalf("Error writing kata JSON file: %v", err)
        os.Exit(1)
    }

    fmt.Printf("✅ Kata created successfully: %s\n", jsonPath)
    fmt.Println(string(out))
}

func loadWords(filepath string) ([]string, error) {
    data, err := os.ReadFile(filepath)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %v", err)
    }
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    return lines, nil
}

func loadTablesFromJSON(filepath string) ([]Table, error) {
    data, err := os.ReadFile(filepath)
    if err != nil {
        return nil, fmt.Errorf("failed to read JSON file: %v", err)
    }
    var tables []Table
    err = json.Unmarshal(data, &tables)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
    }
    return tables, nil
}

func randomChoice(choices []string) string {
    n := sRandInt(len(choices))
    return choices[n]
}

func sRandInt(max int) int {
    if max <= 0 {
        panic("max must be greater than 0")
    }
    var n uint64
    binary.Read(rand.Reader, binary.LittleEndian, &n)
    return int(n % uint64(max))
}

func sShuffle(slice []string) {
    n := len(slice)
    for i := n - 1; i > 0; i-- {
        j := sRandInt(i + 1)
        slice[i], slice[j] = slice[j], slice[i]
    }
}

func addTableWithDependencies(tableMap, selected map[string]Table, name string) {
    if _, exists := selected[name]; exists {
        return
    }
    table, exists := tableMap[name]
    if !exists {
        return
    }
    for _, dep := range table.Dependencies {
        addTableWithDependencies(tableMap, selected, dep)
    }
    selected[name] = table
}
