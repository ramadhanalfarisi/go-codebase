# Services Folder Documentation

## Overview

The `services` folder contains the business logic of your Go application, organized following a **Clean Architecture** pattern. This structure separates concerns into distinct layers, making the codebase maintainable, testable, and scalable.

---

## Table of Contents

1. [Architecture Pattern](#architecture-pattern)
2. [File Structure](#file-structure)
3. [HTTP Methods Explanation](#http-methods-explanation)
4. [Layer-by-Layer Breakdown](#layer-by-layer-breakdown)
5. [Service Examples](#service-examples)

---

## Architecture Pattern

The services folder follows a **Layered Architecture (Clean Architecture)** pattern with the following layers:

```
Request → Controller → Usecase → Repository → Database
          ↓                                      ↓
       (Handle HTTP)      (Business Logic)   (Data Access)
```

### Benefits of This Pattern:

- **Separation of Concerns**: Each layer has a single responsibility
- **Testability**: Layers can be tested independently
- **Maintainability**: Easy to locate and modify functionality
- **Reusability**: Business logic can be reused across different interfaces (REST API, GraphQL)
- **Dependency Inversion**: Higher layers depend on abstractions (interfaces), not concrete implementations

---

## File Structure

```
services/
├── common/                          # Shared utilities across services
│   └── models/
│       └── model.go                # Common model definitions
│
├── product/                         # Product Service
│   ├── controller/
│   │   ├── controller_api.go       # REST API handlers
│   │   ├── controller_graphql.go   # GraphQL handlers
│   │   └── controller_interface.go # Controller interfaces
│   │
│   ├── models/
│   │   └── model.go                # Product data structures
│   │
│   ├── repository/
│   │   ├── repository.go           # Database operations
│   │   └── repository_interface.go # Repository interface
│   │
│   ├── routes/
│   │   ├── api.go                  # REST API route definitions
│   │   └── graphql.go              # GraphQL route definitions
│   │
│   └── usecase/
│       ├── usecase.go              # Business logic
│       └── usecase_interface.go    # Usecase interface
│
└── user/                            # User Service
    ├── controller/
    │   ├── controller.go           # User handlers
    │   └── controller_interface.go # Controller interface
    │
    ├── models/
    │   └── model.go                # User data structures
    │
    ├── repository/
    │   ├── repository.go           # Database operations
    │   └── repository_interface.go # Repository interface
    │
    ├── routes/
    │   └── api.go                  # REST API route definitions
    │
    └── usecase/
        ├── usecase.go              # Business logic
        └── usecase_interface.go    # Usecase interface
```

---

## HTTP Methods Explanation

### **GET** - Retrieve Data

Fetches data from the server without modifying anything.

**Characteristics:**

- Safe operation (no side effects)
- Idempotent (multiple calls return the same result)
- Used to retrieve data
- Parameters passed in URL path or query string
- No request body

**Example in Product Service:**

```go
app.Get("products", controller.GetProducts)           // Get all products
app.Get("products/:id", controller.GetProductById)    // Get product by ID
```

**Use Cases:**

- Fetching a list of all products
- Fetching a specific product by ID
- Fetching user profile information
- Fetching filtered data

---

### **POST** - Create Data

Creates new resource on the server.

**Characteristics:**

- Not idempotent (repeated calls create multiple resources)
- Used to submit data for processing
- Data sent in request body
- Server generates resource identifier
- Returns HTTP 201 (Created) or 200 (OK)

**Example in Product Service:**

```go
app.Post("products", controller.CreateProduct)
```

**Example in User Service:**

```go
app.Post("/register", userController.UserRegister)
app.Post("/login", userController.UserLogin)
```

**Use Cases:**

- Creating a new product
- User registration
- User login
- Creating any new resource

**Implementation:**

```go
func (p *ProductControllerAPI) CreateProduct(c fiber.Ctx) error {
    // 1. Parse JSON from request body
    var productInput models.ProductInput
    err := json.Unmarshal(c.Body(), &productInput)

    // 2. Validate input
    msgs, isValid := helpers.Validate(productInput)

    // 3. Call usecase (business logic)
    product, err := p.usecase.CreateProduct(productInput)

    // 4. Return response
    return succesResponse.SendResponse(c)
}
```

---

### **PUT** - Replace Data Completely

Replaces the entire resource with new data. All fields must be provided.

**Characteristics:**

- Idempotent (repeated calls have the same effect)
- Replaces the entire resource
- All required fields must be provided
- Returns the updated resource
- If resource doesn't exist, some APIs create it

**Example in Product Service:**

```go
app.Put("products/:id", controller.UpdatePutProduct)
```

**Use Cases:**

- Replacing an entire product's information
- Updating user profile with all fields

**Implementation:**

```go
// ProductUpdatePutInput - ALL fields are required
type ProductUpdatePutInput struct {
    Name        string  `json:"name" validate:"required"`
    Price       float64 `json:"price" validate:"required,number"`
    Description string  `json:"description" validate:"required"`
}

func (p *ProductRepository) UpdatePutProduct(id int, input models.ProductUpdatePutInput) (models.Product, error) {
    // Replace all fields at once
    query, args := query_builder.New("product").Update().
        Set("name", input.Name).
        Set("description", input.Description).
        Set("price", input.Price).
        Where("id = ?", id).
        Build("id", "name", "description", "price")

    var product models.Product
    err := p.queryHelper.Update(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
    return product, err
}
```

---

### **PATCH** - Partial Update

Updates only the fields provided, leaving others unchanged.

**Characteristics:**

- Idempotent (repeated calls have the same effect)
- Partially updates a resource
- Only provided fields are updated
- Missing fields are left unchanged
- More efficient than PUT for small updates

**Example in Product Service:**

```go
app.Patch("products/:id", controller.UpdatePatchProduct)
```

**Use Cases:**

- Updating only the product name without touching price or description
- Partial user profile updates
- Updating specific fields without knowing all data

**Implementation:**

```go
// ProductUpdateInput - ALL fields are optional (pointers)
type ProductUpdateInput struct {
    Name        *string  `json:"name"`
    Price       *float64 `json:"price" validate:"omitempty,number"`
    Description *string  `json:"description"`
}

func (p *ProductRepository) UpdateProduct(id int, input models.ProductUpdateInput) (models.Product, error) {
    update := query_builder.New("product").Update()

    // Only set fields that are provided (not nil)
    if input.Name != nil {
        update.Set("name", *input.Name)
    }
    if input.Description != nil {
        update.Set("description", *input.Description)
    }
    if input.Price != nil {
        update.Set("price", *input.Price)
    }

    query, args := update.Where("id = ?", id).Build("id", "name", "description", "price")
    var product models.Product
    err := p.queryHelper.Update(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
    return product, err
}
```

---

### **DELETE** - Remove Data

Removes a resource from the server.

**Characteristics:**

- Idempotent (repeated calls on deleted resource return same result)
- Removes the resource
- Usually identified by ID in URL path
- Returns 204 (No Content) or 200 (OK)

**Example in Product Service:**

```go
app.Delete("products/:id", controller.DeleteProduct)
```

**Use Cases:**

- Deleting a product
- Deleting a user account
- Removing any resource

**Implementation:**

```go
func (p *ProductControllerAPI) DeleteProduct(c fiber.Ctx) error {
    // 1. Get ID from URL parameter
    id := c.Params("id")
    idInt := helpers.StringToInt(id)

    // 2. Call usecase to delete
    prod, err := p.usecase.DeleteProduct(idInt)

    // 3. Return response
    return succesResponse.SendResponse(c)
}
```

---

## Layer-by-Layer Breakdown

### **1. Routes Layer** (`routes/api.go`)

**Purpose**: Defines HTTP route mappings and initializes the dependency injection chain.

**Responsibilities**:

- Define URL endpoints
- Map HTTP methods to controller handlers
- Initialize repository, usecase, and controller instances
- Set up dependency injection

**Example - Product Routes:**

```go
func ProductAPIRoutes(db *sql.DB, app fiber.Router) {
    // Initialize dependencies (Dependency Injection)
    repo := repository.NewProductRepository(db)
    usecase := usecase.NewProductUsecase(repo)
    controller := controller.NewProductControllerAPI(usecase)

    // Define routes and map to handlers
    app.Get("products", controller.GetProducts)
    app.Get("products/:id", controller.GetProductById)
    app.Patch("products/:id", controller.UpdatePatchProduct)
    app.Put("products/:id", controller.UpdatePutProduct)
    app.Post("products", controller.CreateProduct)
    app.Delete("products/:id", controller.DeleteProduct)
}
```

**Flow**: Request comes in → Route matches → Controller handler is called

---

### **2. Controller Layer** (`controller/`)

**Purpose**: Handles HTTP requests/responses and request validation.

**Responsibilities**:

- Parse incoming HTTP request
- Validate request data
- Call usecase for business logic
- Format and return HTTP response
- Handle HTTP status codes

**Files**:

- `controller_interface.go`: Defines controller contract
- `controller_api.go`: REST API implementation
- `controller_graphql.go`: GraphQL implementation

**Example - Product Create Handler:**

```go
func (p *ProductControllerAPI) CreateProduct(c fiber.Ctx) error {
    // Step 1: Parse request body
    var productInput models.ProductInput
    err := json.Unmarshal(c.Body(), &productInput)
    if err != nil {
        return helpers.Response{Code: fiber.StatusBadRequest, ...}.SendResponse(c)
    }

    // Step 2: Validate input
    msgs, isValid := helpers.Validate(productInput)
    if !isValid {
        return helpers.Response{Code: fiber.StatusBadRequest, ...}.SendResponse(c)
    }

    // Step 3: Call business logic
    product, err := p.usecase.CreateProduct(productInput)
    if err != nil {
        return helpers.Response{Code: fiber.StatusInternalServerError, ...}.SendResponse(c)
    }

    // Step 4: Return success response
    return helpers.ResponseData{Code: fiber.StatusOK, Data: product, ...}.SendResponse(c)
}
```

**Key Points**:

- **No business logic** here (that's in usecase)
- Focus on request/response handling
- Always validate input before passing to usecase
- Proper HTTP status codes

---

### **3. Usecase Layer** (`usecase/`)

**Purpose**: Contains the business logic and orchestrates between controller and repository.

**Responsibilities**:

- Implement business rules
- Call repository for data operations
- Perform validation and transformations
- Handle business errors
- Orchestrate complex operations

**Files**:

- `usecase_interface.go`: Defines usecase contract
- `usecase.go`: Business logic implementation

**Example - Product Create Usecase:**

```go
type ProductUsecase struct {
    repository repository.ProductRepositoryInterface
}

func (p *ProductUsecase) CreateProduct(input models.ProductInput) (models.Product, error) {
    // Business logic: Validate and create product
    prod, err := p.repository.CreateProduct(input)
    if err != nil {
        return models.Product{}, errors.New("failed to create product")
    }
    return prod, nil
}
```

**Key Points**:

- **Pure business logic**
- No HTTP knowledge (no fiber.Ctx)
- Depends on repository interface (not concrete implementation)
- Easy to test and reuse

---

### **4. Repository Layer** (`repository/`)

**Purpose**: Handles all data access operations (database queries).

**Responsibilities**:

- Execute SQL queries
- Map database results to models
- Handle database errors
- Provide data access abstraction

**Files**:

- `repository_interface.go`: Defines repository contract
- `repository.go`: Database operations implementation

**Example - Product Repository:**

```go
type ProductRepository struct {
    db          *sql.DB
    queryHelper helpers.QueryHelperInterface
}

func (p *ProductRepository) CreateProduct(input models.ProductInput) (models.Product, error) {
    // Build SQL query
    query, args := query_builder.New("product").
        Insert("name", "description", "price", "created_at").
        Values(input.Name, input.Description, input.Price, time.Now()).
        Build("id", "name", "description", "price")

    // Execute and map result
    var product models.Product
    err := p.queryHelper.Insert(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
    if err != nil {
        return models.Product{}, err
    }
    return product, nil
}
```

**Key Points**:

- **Only data access code** here
- No business logic
- Uses query builder for SQL generation
- Maps database rows to Go structs

---

### **5. Models Layer** (`models/`)

**Purpose**: Defines data structures used across layers.

**File**: `model.go`

**Structure Types**:

**Domain Models** - Used within the application:

```go
type Product struct {
    Id          int     `json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
}
```

**Input Models** - Used for request validation:

```go
// For POST/PUT requests (all fields required)
type ProductInput struct {
    Name        string  `json:"name" validate:"required"`
    Price       float64 `json:"price" validate:"required,number"`
    Description string  `json:"description" validate:"required"`
}

// For PATCH requests (all fields optional)
type ProductUpdateInput struct {
    Name        *string  `json:"name"`
    Price       *float64 `json:"price" validate:"omitempty,number"`
    Description *string  `json:"description"`
}

// For PUT requests (all fields required)
type ProductUpdatePutInput struct {
    Name        string  `json:"name" validate:"required"`
    Price       float64 `json:"price" validate:"required,number"`
    Description string  `json:"description" validate:"required"`
}
```

**Key Points**:

- Use `*T` (pointer) for optional fields (PATCH)
- Use `T` (value) for required fields (POST/PUT)
- Include validation tags: `validate:"required"`, `validate:"omitempty,number"`
- JSON tags for serialization: `json:"fieldname"`

---

## Service Examples

### Example 1: Product Service - Complete Flow

**Request**: `GET /products/1`

```
1. HTTP Request arrives
   ↓
2. routes/api.go matches route
   → app.Get("products/:id", controller.GetProductById)
   ↓
3. controller_api.go - GetProductById handler
   → Extract ID from URL parameter
   → Validate ID exists
   → Call usecase.GetProductById(idInt)
   ↓
4. usecase.go - GetProductById
   → Call repository.GetProductById(id)
   → Handle errors
   ↓
5. repository.go - GetProductById
   → Build SQL: SELECT id, name, description, price FROM product WHERE id = ?
   → Execute query
   → Map result to Product struct
   ↓
6. Return through layers (Repository → Usecase → Controller)
   ↓
7. Controller formats response and sends HTTP response
   ↓
8. Client receives JSON with product data
```

**Response**:

```json
{
    "code": 200,
    "status": constants.StatusSuccess,
    "message": "Product retrieved successfully",
    "data": {
        "id": 1,
        "name": "Laptop",
        "price": 999.99,
        "description": "High-performance laptop"
    }
}
```

---

### Example 2: Product Service - Create with Validation

**Request**: `POST /products` with body:

```json
{
  "name": "New Product",
  "price": 49.99,
  "description": "Amazing product"
}
```

```
1. HTTP Request arrives with JSON body
   ↓
2. routes/api.go matches route
   → app.Post("products", controller.CreateProduct)
   ↓
3. controller_api.go - CreateProduct handler
   → Parse JSON: json.Unmarshal(c.Body(), &productInput)
   → Validate: helpers.Validate(productInput)
      - Checks: required fields, data types
   → Call usecase.CreateProduct(productInput)
   ↓
4. usecase.go - CreateProduct
   → Call repository.CreateProduct(input)
   ↓
5. repository.go - CreateProduct
   → Build SQL INSERT query
   → Execute query: INSERT INTO product (name, description, price, created_at) VALUES (...)
   → Return inserted product with generated ID
   ↓
6. Return through layers
   ↓
7. Controller returns HTTP 200 with created product
```

**Response**:

```json
{
    "code": 200,
    "status": constants.StatusSuccess,
    "message": "Product created successfully",
    "data": {
        "id": 5,
        "name": "New Product",
        "price": 49.99,
        "description": "Amazing product"
    }
}
```

---

### Example 3: Difference Between PATCH and PUT

**PATCH Request**: `PATCH /products/1` with body:

```json
{
  "price": 79.99
}
```

- Only updates `price` field
- `name` and `description` remain unchanged
- Uses `ProductUpdateInput` (all fields optional with pointers)

**PUT Request**: `PUT /products/1` with body:

```json
{
  "name": "Updated Product",
  "price": 99.99,
  "description": "Updated description"
}
```

- Replaces all fields
- Missing any field causes validation error
- Uses `ProductUpdatePutInput` (all fields required)

---

### Example 4: User Service - Login Flow

**Request**: `POST /login` with body:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

```
1. HTTP Request with login credentials
   ↓
2. routes/api.go matches route
   → app.Post("/login", userController.UserLogin)
   ↓
3. controller.go - UserLogin
   → Parse JSON: json.Unmarshal(c.Body(), &userLoginInput)
   → Validate input
   → Call usecase.UserLogin(userLoginInput)
   ↓
4. usecase.go - UserLogin
   → Call repository.GetUserByEmail(input)
   → Verify password matches
   → Generate JWT token
   ↓
5. repository.go - GetUserByEmail
   → Build SQL: SELECT id, email, password, roles FROM users WHERE email = ?
   → Execute query
   ↓
6. Return through layers
   ↓
7. Controller returns HTTP 200 with JWT token
```

**Response**:

```json
{
    "code": 200,
    "status": constants.StatusSuccess,
    "message": "Login successful",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
}
```

---

## Summary Table

| Layer          | File(s)             | Purpose          | Contains                                         |
| -------------- | ------------------- | ---------------- | ------------------------------------------------ |
| **Routes**     | `routes/api.go`     | Endpoint mapping | Route definitions, Dependency Injection          |
| **Controller** | `controller_api.go` | HTTP handling    | Request parsing, Validation, Response formatting |
| **Usecase**    | `usecase.go`        | Business logic   | Business rules, Orchestration, Error handling    |
| **Repository** | `repository.go`     | Data access      | SQL queries, Database operations                 |
| **Models**     | `model.go`          | Data structures  | Domain models, Input models, DTOs                |

---

## Best Practices

1. **Always use interfaces** (`*_interface.go` files) for dependency injection
2. **Keep controllers thin** - only handle HTTP concerns
3. **Put all business logic in usecase** - keeps it testable and reusable
4. **Use repository for all database operations** - allows easy testing with mocks
5. **Validate early** - in controller layer before passing to usecase
6. **Use pointers for optional fields** in input models for PATCH requests
7. **Handle errors at each layer** and provide meaningful error messages
8. **Separate POST/PUT input models** - POST for creation, PUT for full updates, PATCH for partial updates

---

## Common Patterns

### Creating a New Service

1. Create folder structure: `myservice/controller/`, `usecase/`, `repository/`, `models/`, `routes/`
2. Define models in `models/model.go`
3. Create interfaces in `*_interface.go` files
4. Implement repository methods for database operations
5. Implement usecase for business logic
6. Implement controller for HTTP handling
7. Define routes in `routes/api.go`
8. Register routes in main application

### Testing Strategy

- **Repository**: Mock database, test SQL generation
- **Usecase**: Mock repository, test business logic
- **Controller**: Mock usecase, test HTTP request/response handling
