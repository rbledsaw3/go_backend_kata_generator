# Kata Generator Project

This repository provides a structured workflow to quickly spin up mini-projects (katas) designed for practicing backend development with Go, Goose, sqlc, and PostgreSQL.

## 📁 Project Structure

```
.
├── kata_generator
│   ├── adjectives.txt          # Adjectives for project naming
│   ├── nouns.txt               # Nouns for project naming
│   ├── schema_templates.json   # JSON schema definitions for tables
│   ├── go.mod                  # Go module file
│   └── main.go                 # Kata generation script
├── katas
│   └── <project_name>
│       └── <project_name>.json # Generated kata specification
└── kata_tester
    └── go.mod                  # Future testing suite (to be implemented)
```

## 🚀 Getting Started

### Step 1: Kata Generation

Navigate to the `kata_generator` directory:

```bash
cd kata_generator
```

Run the generator:

```bash
go run main.go
```

This command:
- Generates a random project name.
- Creates a schema from predefined table templates, resolving all dependencies.
- Outputs a JSON file under `../katas/<project_name>/<project_name>.json`.

Example generated kata:

```json
{
  "project_name": "silver-lake",
  "tables": {
    "users": ["id uuid primary key", "username text unique", "email text", "created_at timestamp"],
    "products": ["id uuid primary key", "name text", "price integer", "stock integer"],
    "orders": ["id uuid primary key", "user_id uuid references users(id)", "product_id uuid references products(id)", "quantity integer", "created_at timestamp"]
  }
}
```

### Step 2: Manual Implementation

Based on the generated kata JSON, you manually:
- Set up Goose migration files (`.sql`) to create the schema.
- Configure `sqlc` to generate the corresponding database interaction code.
- Write RESTful API endpoints in Go to interact with the database.

### Step 3: Automated Testing (Future)

The `kata_tester` module (under construction) will:
- Read the generated kata JSON specification.
- Run automated HTTP request tests against your implemented API endpoints.
- Verify correct database read/write operations.

## 🛠️ Technologies

- **Go**: Main programming language for API implementation.
- **Goose**: Database migration management.
- **sqlc**: SQL code generation tool.
- **PostgreSQL**: Relational database system.

## 🔮 Upcoming Features

- Complete implementation of `kata_tester` for automated API and database integrity checks.

## 📌 Notes

- Ensure your PostgreSQL environment is set up and running.
- Update `adjectives.txt`, `nouns.txt`, and `schema_templates.json` to expand vocabulary and schema complexity as desired.

## ✏️ Extending Schema Templates

You can easily extend your schema templates by adding entries in `schema_templates.json`. Here's an example:

```json
[
  {
    "name": "categories",
    "columns": [
      "id uuid primary key",
      "name text unique",
      "description text"
    ],
    "dependencies": []
  },
  {
    "name": "tags",
    "columns": [
      "id uuid primary key",
      "name text unique"
    ],
    "dependencies": []
  },
  {
    "name": "product_tags",
    "columns": [
      "product_id uuid references products(id)",
      "tag_id uuid references tags(id)",
      "primary key (product_id, tag_id)"
    ],
    "dependencies": ["products", "tags"]
  }
]
```

---

Enjoy refining your backend skills through structured practice katas!
