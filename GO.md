# Go Essentials - Guía Completa

## Tabla de Contenidos
1. [Estructura Básica](#estructura-básica)
2. [Módulos y Packages](#módulos-y-packages)
3. [Variables y Constantes](#variables-y-constantes)
4. [Tipos de Datos](#tipos-de-datos)
5. [Operadores](#operadores)
6. [Control de Flujo](#control-de-flujo)
7. [Manejo de Errores](#manejo-de-errores)
8. [Punteros](#punteros)
9. [Arrays, Slices y Maps](#arrays-slices-y-maps)
10. [Structs](#structs)
11. [Métodos y Receivers](#métodos-y-receivers)
12. [Custom Types](#custom-types)
13. [Funciones como Valores](#funciones-como-valores)
14. [Funciones Anónimas y Closures](#funciones-anónimas-y-closures)
15. [Funciones Variádicas](#funciones-variádicas)
16. [Recursión](#recursión)
17. [Interfaces](#interfaces)
18. [Generics](#generics)
19. [Concurrencia (Goroutines)](#concurrencia-goroutines)
20. [REST API con Gin](#rest-api-con-gin)
21. [Bases de Datos con `database/sql`](#bases-de-datos-con-databasesql)
22. [Recursos Útiles](#recursos-útiles)

---

## Estructura Básica

Toda aplicación Go requiere:
1. **Package**: Identificar el paquete (convención: mismo nombre que el directorio)
2. **Imports**: Importar las dependencias necesarias
3. **Código**: La lógica de la aplicación

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

**Ejecución de un archivo:**
```bash
go run archivo.go
```

**Convenciones de nombres:**
- Variables y funciones: `camelCase` (ej: `investmentAmount`, `getUserData`)
- Constantes: `UPPER_SNAKE_CASE` o `camelCase`
- Archivos: `snake_case` (ej: `hello_world.go`)

---

## Módulos y Packages

### Jerarquía
Un módulo contiene packages, y los packages contienen archivos.

```
Módulo
  └─ Package 1
       └─ archivo.go
  └─ Package 2
       └─ archivo.go
```

### Iniciar un módulo
```bash
go mod init example.com/project-name
```

El archivo `go.mod` se genera automáticamente:
```go
module example.com/investment-calculator

go 1.21.2
```

Las URLs usan guiones (`-`), no guiones bajos (`_`).

### Ejecutar un módulo
```bash
go run .
```

Ejecuta automáticamente el `main` package con la función `main()`.

### Construir un ejecutable
```bash
go build
```

Genera un ejecutable en el directorio actual.

### Package `main` vs Librerías
- **`main`**: Nombre reservado. Solo usado en programas ejecutables.
  - Requiere una función `main()` como punto de entrada.
- **Otros packages**: Usados para librerías o módulos reutilizables.
  - No requieren `main()` porque no se ejecutan directamente.

---

## Variables y Constantes

### Variables
Se declaran con `var`:

```go
var investmentAmount = 1000      // tipo inferido
var investmentAmount int = 1000  // tipo explícito
investmentAmount := 1000         // declaración corta (solo en funciones)
```

La declaración corta (`:=`) es más común por ser concisa.

**Inicialización múltiple:**
```go
var firstName, lastName string = "John", "Doe"
firstName, lastName := "John", "Doe"
```

### Constantes
Se declaran con `const`:

```go
const inflationRate = 2.5
const maxAttempts int = 5
```

Las constantes **no pueden cambiar** y deben conocerse en tiempo de compilación.

---

## Tipos de Datos

### Tipos Primitivos

| Tipo | Valor Nulo | Descripción |
|------|-----------|-------------|
| `int` | 0 | Entero (tamaño depende de la plataforma) |
| `int8`, `int16`, `int32`, `int64` | 0 | Enteros con tamaño específico |
| `uint` | 0 | Entero sin signo (positivos) |
| `uint8`, `uint16`, `uint32`, `uint64` | 0 | Sin signo con tamaño específico |
| `float32` | 0.0 | Punto flotante de 32 bits |
| `float64` | 0.0 | Punto flotante de 64 bits |
| `string` | `""` | Texto |
| `bool` | `false` | Booleano |
| `byte` | 0 | Alias de `uint8` |
| `rune` | 0 | Alias de `int32`, representa un carácter Unicode |

### Conversión de Tipos

```go
age := 32
ageFloat := float64(age)

text := "123"
number, err := strconv.ParseFloat(text, 64)
```

### Operadores Aritméticos

```go
a := 10
b := 3

sum := a + b        // 13
diff := a - b       // 7
product := a * b    // 30
quotient := a / b   // 3
remainder := a % b  // 1

// Operadores compuestos
a += 5              // a = a + 5
a -= 2              // a = a - 2
a *= 3              // a = a * 3
a /= 2              // a = a / 2
```

### Operadores Lógicos

```go
a := true
b := false

a && b  // AND (ambos true)
a || b  // OR (al menos uno true)
!a      // NOT (inverso)

// Comparación
x := 10
y := 5

x == y  // igual
x != y  // diferente
x > y   // mayor
x < y   // menor
x >= y  // mayor o igual
x <= y  // menor o igual
```

---

## Control de Flujo

### Condicionales: `if`, `else if`, `else`

```go
age := 25

if age >= 18 {
    fmt.Println("Es adulto")
} else if age >= 13 {
    fmt.Println("Es adolescente")
} else {
    fmt.Println("Es menor")
}
```

### Switch Statement

```go
choice := 2

switch choice {
case 1:
    fmt.Println("Opción 1")
case 2:
    fmt.Println("Opción 2")
case 3:
    fmt.Println("Opción 3")
default:
    fmt.Println("Opción inválida")
}
```

### Loops: `for`

Go **solo tiene un tipo de loop:** `for`

**Loop tradicional:**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)  // 0, 1, 2, 3, 4
}
```

**Loop while:**
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

**Loop infinito (con break):**
```go
for {
    fmt.Println("Presiona Ctrl+C para salir")
    break  // o continue
}
```

**Loop sobre rango (arrays, slices, strings):**
```go
numbers := []int{1, 2, 3, 4, 5}

for index, value := range numbers {
    fmt.Printf("%d: %v\n", index, value)
}

// Si solo necesitas el índice
for i := range numbers {
    fmt.Println(i)
}

// Si solo necesitas el valor
for _, v := range numbers {
    fmt.Println(v)
}
```



## Manejo de Errores

Go **no usa excepciones**. En su lugar, las funciones que podrían fallar retornan un valor `error` junto con el resultado.

### Patrón de Error
La convención es devolver el resultado y luego el error:

```go
import (
    "os"
    "strconv"
    "errors"
)

func readFromFile(filename string) (float64, error) {
    data, err := os.ReadFile(filename)
    
    if err != nil {
        return 0, errors.New("failed to read file")
    }
    
    valueText := string(data)
    value, err := strconv.ParseFloat(valueText, 64)
    
    if err != nil {
        return 0, errors.New("failed to parse value")
    }
    
    return value, nil
}

// Uso
balance, err := readFromFile("balance.txt")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Balance:", balance)
```

**Ventaja:** Los errores no rompen la aplicación; el programador decide qué hacer con ellos.

### Crear errores personalizados
```go
if firstName == "" {
    return nil, errors.New("first name is required")
}
```

---

## Punteros

Un **puntero** es una variable que almacena la dirección de memoria de otra variable.

### Declaración y Operadores
```go
var age int = 32
var agePointer *int  // declaración sin valor
agePointer = &age    // & obtiene la dirección de memory

// Acceso
fmt.Println(age)         // 32
fmt.Println(agePointer)  // 0xc00000a088 (dirección de memoria)
fmt.Println(*agePointer) // 32 (desreferencia: obtiene el valor)
```

| Operador | Uso | Descripción |
|----------|-----|-------------|
| `&` | `&variable` | Obtiene la dirección de memoria |
| `*` | `*pointer` | Desreferencia: obtiene el valor apuntado |
| `*` | `var p *int` | Declara un puntero a `int` |

### Paso de Punteros a Funciones
```go
func addAdultYears(age *int) {
    *age = *age + 18  // desreferencia para modificar
}

func main() {
    age := 20
    addAdultYears(&age)
    fmt.Println(age)  // 38 (modificado)
}
```

**Cuando el puntero apunta a una variable:**
```go
age := 20
agePointer := &age

// Ambas formas son equivalentes
addAdultYears(&age)
addAdultYears(agePointer)
```

### Cuándo usar Punteros
- **No siempre:** Modificar valores primitivos (`int`, `string`) es innecesario.
- **Sí:** Modificar estructuras complejas (`struct`) desde funciones.
- **Sí:** Pasar estructuras grandes sin copiar (eficiencia de memoria).

---

## Arrays, Slices y Maps

### Arrays

Un **array** es una colección de elementos **del mismo tipo con tamaño fijo**.

**Definición:**
```go
var numbers [5]int              // array de 5 enteros (valores por defecto: 0)
var names [3]string             // array de 3 strings (valores por defecto: "")
prices := [4]float64{10, 20, 30, 40}  // inicialización con valores
```

**Acceso:**
```go
prices := [4]float64{10, 20, 30, 40}

fmt.Println(prices[0])    // 10 (primer elemento)
prices[1] = 25            // modificar elemento
fmt.Println(len(prices))  // 4 (tamaño del array)
```

**Limitaciones:**
- El tamaño es **fijo** y debe conocerse en compilación.
- No se puede agregar ni eliminar elementos.
- No se puede cambiar el tamaño.

### Slices

Un **slice** es una **vista dinámica sobre un array**. Es más flexible que un array.

**Creación de Slices:**

```go
// Slice sin inicialización (vacío)
var prices []float64

// Slice con valores
hobbies := []string{"drawing", "guitar", "anime"}

// Slice a partir de un array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // elementos del índice 1 a 3 (4 no incluido)
```

**Sintaxis de Slicing:**
```go
arr := [5]int{10, 20, 30, 40, 50}

arr[1:4]   // {20, 30, 40} - del índice 1 al 3
arr[:3]    // {10, 20, 30} - desde el principio hasta el índice 2
arr[2:]    // {30, 40, 50} - desde el índice 2 hasta el final
arr[:]     // {10, 20, 30, 40, 50} - todo el array
```

**Importante:** El último índice **no se incluye** (es exclusivo).

### Slices: Len y Cap

```go
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice := arr[2:5]  // {3, 4, 5}

fmt.Println(len(slice))  // 3 (cantidad de elementos)
fmt.Println(cap(slice))  // 8 (capacidad disponible desde el inicio de slice)
```

- **`len`**: Cantidad actual de elementos en el slice.
- **`cap`**: Elementos disponibles desde donde comienza el slice hasta el final del array subyacente.

### Slices: Mutabilidad

Los slices **apuntan al array original**. Cambios en el slice afectan el array:

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4]  // {20, 30, 40}

slice[0] = 999     // modifica arr[1]
fmt.Println(arr)   // [10, 999, 30, 40, 50]
```

### Slices: Append

La función `append` agrega elementos a un slice, pero **devuelve un nuevo slice**.

```go
prices := []float64{10.99, 8.99}
prices = append(prices, 12.99)       // se reasigna al nuevo slice

fmt.Println(prices)  // [10.99, 8.99, 12.99]
```

**Agregar múltiples elementos:**
```go
original := []int{1, 2, 3}
toAdd := []int{4, 5, 6}

// El operador ... "desempaqueta" el slice
original = append(original, toAdd...)  // [1, 2, 3, 4, 5, 6]
```

### Slices: Eliminar Elementos

Para eliminar elementos, usa slicing:

```go
prices := []float64{10.99, 8.99, 12.99, 5.99}

// Eliminar el primer elemento
prices = prices[1:]  // [8.99, 12.99, 5.99]

// Eliminar el último elemento
prices = prices[:len(prices)-1]  // [8.99, 12.99]

// Eliminar elemento en posición 1
prices = append(prices[:1], prices[2:]...)  // [8.99]
```

### Slices: Make

La función `make` crea un slice con longitud y capacidad predefinidas:

```go
// make(tipo, longitud, capacidad)
numbers := make([]int, 5, 10)  // slice de 5 elementos con capacidad 10

fmt.Println(len(numbers))  // 5 (cantidad de elementos actuales)
fmt.Println(cap(numbers))  // 10 (espacio total disponible)
```

**Parámetros:**
- **Primer argumento:** Tipo del slice.
- **Segundo argumento:** Longitud inicial (`len`).
- **Tercer argumento (opcional):** Capacidad máxima (`cap`).

**Inicialmente con valores por defecto:**
```go
numbers := make([]int, 5)  // [0, 0, 0, 0, 0]

// Acceso
numbers[0] = 10
fmt.Println(numbers)  // [10, 0, 0, 0, 0]
```

**Agregar elementos sin reasignación:**
```go
numbers := make([]int, 2, 5)  // longitud 2, capacidad 5
numbers[0] = 1
numbers[1] = 2

// append usa la capacidad disponible sin reasignar
numbers = append(numbers, 3, 4)  // ahora tiene 4 elementos
fmt.Println(len(numbers), cap(numbers))  // 4 5

// Al exceder la capacidad, Go la duplica
numbers = append(numbers, 5)
fmt.Println(len(numbers), cap(numbers))  // 5 10 (capacidad se duplicó)
```

**Ventajas de `make`:**
- Evita reasignaciones múltiples si sabes aproximadamente cuántos elementos necesitas.
- Mejor performance cuando trabajas con colecciones grandes.
- Más eficiente de memoria.

**Cuándo usar `make`:**
- Si sabes la cantidad de elementos (usa `make`).
- Si es un número pequeño o desconocido, usa inicialización directa: `[]int{1, 2, 3}`.

### Maps

Un **map** es una colección de pares clave-valor (similar a diccionarios u objetos).

**Creación:**
```go
// Map vacío
var user map[string]string

// Map con valores iniciales
person := map[string]string{
    "name": "John",
    "age":  "30",
}

// Map con make (con capacidad inicial)
scores := make(map[string]int)
userAges := make(map[string]int, 10)  // capacidad inicial de 10 pares
```

**Acceso y Modificación:**
```go
person := map[string]string{
    "name": "John",
    "city": "New York",
}

fmt.Println(person["name"])  // "John"

person["age"] = "30"         // agregar par clave-valor
person["name"] = "Jane"      // modificar valor

delete(person, "city")       // eliminar una entrada
```

**Verificar si existe una clave:**
```go
person := map[string]string{"name": "John"}

value, exists := person["name"]  // ("John", true)
value, exists := person["age"]   // ("", false)

if exists {
    fmt.Println("Clave encontrada:", value)
}

// Sintaxis más corta
if value, ok := person["name"]; ok {
    fmt.Println(value)
}
```

**Iterar sobre un map:**
```go
person := map[string]string{
    "name": "John",
    "city": "NYC",
}

for key, value := range person {
    fmt.Printf("%s: %s\n", key, value)
}

// Solo claves
for key := range person {
    fmt.Println(key)
}

// Solo valores
for _, value := range person {
    fmt.Println(value)
}
```

**Tipos de claves comunes:**
- `string`: La más común.
- `int`, `float64`: Tipos numéricos.
- Cualquier tipo que tenga comparación de igualdad.

**No permitidas como claves:**
- Slices
- Maps
- Funciones

**Nota sobre orden:** Los maps en Go **no tienen orden garantizado**. Al iterar, el orden es aleatorio.

```go
person := map[string]string{"name": "John", "city": "NYC", "age": "30"}

for k := range person {
    fmt.Println(k)  // orden aleatorio cada ejecución
}
```

### Comparación: Arrays vs Slices vs Maps

| Característica | Array | Slice | Map |
|---|---|---|---|
| Tamaño | Fijo | Dinámico | Dinámico |
| Inicialización | `[5]int{1,2,3}` | `[]int{1,2,3}` | `map[string]int{"a": 1}` |
| Acceso | Índice | Índice | Clave |
| Mutabilidad | Sí | Sí | Sí |
| Valores por defecto | Sí | Sí | No |

---

## Structs

Un **struct** es una colección de campos que forman una estructura de datos.

### Definición
```go
type User struct {
    firstName string    // privado (empieza en minúscula)
    lastName  string
    birthdate string
    createdAt time.Time
}

type Admin struct {
    email    string
    password string
}
```

### Creación de Instancias

**Usando Struct Literal (Composite Literal):**
```go
user := User{
    firstName: "John",
    lastName:  "Doe",
    birthdate: "01/01/1990",
    createdAt: time.Now(),
}
```

**Acceso a campos:**
```go
fmt.Println(user.firstName)  // "John"
```

### Punteros a Structs
```go
var userPointer *User
userPointer = &user

// Formas equivalentes (Go permite la forma corta):
fmt.Println((*userPointer).firstName)  // Técnicamente correcto
fmt.Println(userPointer.firstName)     // Forma corta (recomendada)
```

Go **automáticamente desreferencia** punteros a structs al acceder a campos.

### Struct Tags

Los **tags** son metadatos en los campos de un struct que se usan para serialización, validación, etc.

```go
type Note struct {
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
```

**Tag `json`:** Especifica cómo serializar/deserializar a JSON:
- `json:"fieldName"`: Usa ese nombre en JSON.
- `json:"fieldName,omitempty"`: Omite el campo si está vacío.
- `json:"-"`: Ignora este campo en JSON.

**Ejemplo:**
```go
type User struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Password  string `json:"-"` // nunca se serializa
    Age       int    `json:"age,omitempty"`
}

user := User{
    FirstName: "John",
    LastName:  "Doe",
    Password:  "secret",  // no aparecerá en JSON
    Age:       0,         // omitempty lo excluye si es cero
}

// Serializar a JSON
data, err := json.Marshal(user)
// {"firstName":"John","lastName":"Doe"}
// Nota: Password no aparece, Age tampoco (valor cero)
```

**Otros tags comunes:**
- `xml:"field"`: Para serialización XML.
- `db:"field"`: Para mapeo con bases de datos.
- `validate:"required"`: Para validación (librerías externas).

---

| Campo | Visibilidad | Descripción |
|-------|------------|-------------|
| `firstName` | Privado | Solo accesible dentro del mismo package |
| `FirstName` | Público | Accesible desde cualquier package |

```go
// En user/user.go (package user)
type User struct {
    firstName string  // privado
    FirstName string  // público (incorrecto, rompe encapsulación)
}

// En main.go (package main)
user.firstName   // Error: campo privado
user.FirstName   // OK: campo público
```

**Convención:** Usar campos privados con métodos públicos (`New()`, `OutputDetails()`).

---

## Métodos y Receivers

Un **método** es una función asociada a un tipo específico mediante un **receiver**.

### Sintaxis
```go
func (receiver receiverType) methodName() {
    // código
}
```

### Receptor por Valor vs Puntero

**Receptor por valor (copia):**
```go
func (u User) outputDetails() {
    fmt.Println(u.firstName)  // lee los datos
}

user.outputDetails()  // funciona tanto con User como con *User
```

**Receptor por puntero (modifica original):**
```go
func (u *User) clear() {
    u.firstName = ""  // desreferencia automática
    u.lastName = ""
}

user := User{firstName: "John", lastName: "Doe"}
user.clear()  // Go automáticamente pasa &user

// Después: user.firstName = "", user.lastName = ""
```

**Regla de oro:**
- Receptor por **valor** (`u User`): Para métodos que leen datos.
- Receptor por **puntero** (`u *User`): Para métodos que modifican datos.

### Constructores

En Go no existen constructores nativos. Por convención, se usa una función `New`:

```go
func New(firstName, lastName, birthdate string) (*User, error) {
    if firstName == "" || lastName == "" || birthdate == "" {
        return nil, errors.New("all fields are required")
    }
    
    return &User{
        firstName: firstName,
        lastName:  lastName,
        birthdate: birthdate,
        createdAt: time.Now(),
    }, nil
}

// Uso
user, err := New("John", "Doe", "01/01/1990")
if err != nil {
    fmt.Println("Error:", err)
    return
}
user.OutputDetails()
```

---

## Struct Embedding (Composición)

**Struct Embedding** permite reutilizar un struct dentro de otro sin darle un nombre específico.

### Ejemplo
```go
type User struct {
    firstName string
    lastName  string
    createdAt time.Time
}

type Admin struct {
    email    string
    password string
    User     // embedded struct (anonymous field)
}

func NewAdmin(email, password string) Admin {
    return Admin{
        email:    email,
        password: password,
        User: User{
            firstName: "ADMIN",
            lastName:  "ADMIN",
            createdAt: time.Now(),
        },
    }
}

admin := NewAdmin("admin@example.com", "secret")

// Acceso directo a campos y métodos de User
fmt.Println(admin.firstName)   // "ADMIN"
admin.OutputDetails()          // método heredado de User
```

### Ventajas
- **Sin repetición:** `Admin` hereda todos los campos y métodos de `User`.
- **Composición sobre herencia:** Reutiliza comportamiento sin usar herencia.

---

## Custom Types

Un **Custom Type** es un nuevo tipo basado en un tipo primitivo, con métodos propios.

### Definición
```go
type str string  // nuevo tipo basado en string

func (s str) log() {
    fmt.Println(s)
}

func main() {
    text := str("Hello")
    text.log()  // imprime: Hello
}
```

### Basado en Tipos Primitivos

```go
type Age int
type Temperature float64
type EmailAddress string

func main() {
    myAge := Age(25)
    temp := Temperature(98.6)
    email := EmailAddress("john@example.com")
    
    fmt.Println(myAge, temp, email)
}
```

### Custom Type con Métodos

```go
type Revenue float64
type Expenses float64

func (r Revenue) String() string {
    return fmt.Sprintf("$%.2f", r)
}

func (e Expenses) String() string {
    return fmt.Sprintf("$%.2f", e)
}

func calculateProfit(revenue Revenue, expenses Expenses) float64 {
    return float64(revenue) - float64(expenses)
}

func main() {
    revenue := Revenue(1000.50)
    expenses := Expenses(500.25)
    profit := calculateProfit(revenue, expenses)
    
    fmt.Println("Revenue:", revenue)      // $1000.50
    fmt.Println("Expenses:", expenses)    // $500.25
    fmt.Println("Profit:", profit)        // 500.25
}
```

### Alias vs Tipos Nuevos

```go
// Alias: mismos tipos, solo otro nombre
type strAlias = string

// Tipo nuevo: distintos tipos, no intercambiables
type strType string

// Ambos se inicializan igual
var a strAlias = "hello"
var b strType = "world"

// Pero en funciones es diferente
func printString(s string) {
    fmt.Println(s)
}

printString(a)  // OK (strAlias es string)
printString(b)  // Error: strType no es string
```

### Casos de Uso

- **Agregar métodos a tipos primitivos:** Sin esto, `int` no puede tener métodos.
- **Crear tipos especializados sin structs:** Más simple que definir un struct.
- **Distinción de tipos:** `Revenue` y `Expenses` son diferentes tipos aunque ambos sean `float64`.

---

## Funciones como Valores

En Go, las funciones son **ciudadanos de primera clase**. Pueden ser asignadas a variables, pasadas como argumentos y retornadas desde otras funciones.

### Tipo de Función

Para mayor claridad, se puede definir un **type alias** para funciones:

```go
// Define un tipo de función que toma un int y retorna un int
type transformFunc func(int) int

// Definir función que se ajusta al tipo
func double(n int) int {
    return n * 2
}

func main() {
    numbers := []int{1, 2, 3, 4}
    
    // Pasar función como argumento (sin paréntesis)
    doubled := transformNumbers(&numbers, double)
    fmt.Println(doubled)  // [2, 4, 6, 8]
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
    dNumbers := []int{}
    for _, v := range *numbers {
        dNumbers = append(dNumbers, transform(v))
    }
    return dNumbers
}
```

**Nota importante:** Al pasar una función como argumento, **no se usan paréntesis** (`double`, no `double()`).

### Retornar Funciones

Las funciones pueden retornar otras funciones:

```go
// Factory function: retorna una función
func getTransformerFunction(numbers *[]int) transformFunc {
    if (*numbers)[0] == 1 {
        return double  // retorna sin paréntesis
    } else {
        return triple
    }
}

func double(n int) int {
    return n * 2
}

func triple(n int) int {
    return n * 3
}

// Uso
transformerFn := getTransformerFunction(&numbers)
result := transformNumbers(&numbers, transformerFn)
```

---

## Funciones Anónimas y Closures

Una **función anónima** (lambdas o funciones literales) es una función sin nombre que se define en el lugar donde se usa.

### Sintaxis

```go
func main() {
    numbers := []int{1, 2, 3}
    
    // Función anónima definida inline
    transformed := transformNumbers(&numbers, func(number int) int {
        return number * 2
    })
    fmt.Println(transformed)  // [2, 4, 6]
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
    dNumbers := []int{}
    for _, v := range *numbers {
        dNumbers = append(dNumbers, transform(v))
    }
    return dNumbers
}
```

### Closures (Cierres)

Una **closure** es una función anónima que captura variables del ámbito donde fue creada.

```go
func createTransformer(factor int) func(int) int {
    // La función anónima retornada "captura" la variable factor
    return func(number int) int {
        return number * factor  // factor viene del ámbito externo
    }
}

func main() {
    double := createTransformer(2)    // factor = 2
    triple := createTransformer(3)    // factor = 3
    
    fmt.Println(double(5))  // 10 (5 * 2)
    fmt.Println(triple(5))  // 15 (5 * 3)
}
```

**Cómo funciona:** Cada closure tiene su propia copia de las variables capturadas. Cambios en una closure no afectan a otras.

### Closures con Referencias

Si una closure captura una variable **por referencia**, puede modificarla:

```go
func main() {
    count := 0
    
    increment := func() {
        count++  // modifica la variable externa
    }
    
    fmt.Println(count)  // 0
    increment()
    increment()
    fmt.Println(count)  // 2
}
```

---

## Funciones Variádicas

Una **función variádica** acepta un número variable de argumentos del mismo tipo.

### Sintaxis

```go
// ...type indica que acepta 0 o más argumentos de ese tipo
func sumup(nums ...int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return sum
}

func main() {
    result1 := sumup(1, 10, 15, 40, -5)  // 61
    result2 := sumup(5, 10)               // 15
    result3 := sumup()                    // 0 (sin argumentos)
    
    fmt.Println(result1, result2, result3)
}
```

### Desempaquetar Slices

Para pasar un slice a una función variádica, usa `...` al final del slice:

```go
numbers := []int{1, 10, 15}
result := sumup(numbers...)  // desempaqueta [1, 10, 15] en argumentos individuales
fmt.Println(result)          // 26
```

### Parámetros Fijos + Variádicos

Los parámetros variádicos **deben ir al final**:

```go
func printInfo(prefix string, values ...string) {
    fmt.Print(prefix)
    for _, v := range values {
        fmt.Print(" " + v)
    }
    fmt.Println()
}

printInfo("Items:", "apple", "banana", "cherry")
// Output: Items: apple banana cherry
```

---

## Recursión

**Recursión** es cuando una función se llama a sí misma.

### Ejemplo Básico: Factorial

```go
func noRecursiveFactorial(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result = result * i
    }
    return result
}

func recursiveFactorial(n int) int {
    if n == 0 {  // caso base: evita recursión infinita
        return 1
    }
    return n * recursiveFactorial(n-1)  // llamada recursiva
}

func main() {
    fmt.Println(recursiveFactorial(5))  // 120 (5 * 4 * 3 * 2 * 1)
}
```

**Elementos clave:**
1. **Caso base:** Condición que detiene la recursión (`if n == 0`).
2. **Caso recursivo:** La función se llama con un argumento más pequeño.

### Cuándo Usar Recursión

- **Árboles y grafos:** Recorrido de estructuras jerárquicas.
- **Divide y conquista:** Problemas que se dividen en subproblemas idénticos.
- **Backtracking:** Explorar todas las posibilidades.

**Ventajas:**
- Código elegante y conciso para ciertos problemas.

**Desventajas:**
- Más lenta que iteración (overhead de llamadas).
- Riesgo de stack overflow con valores grandes.

---

## Interfaces

Una **interface** define un conjunto de métodos sin implementarlos. Cualquier tipo que implemente todos esos métodos **satisface** la interface.

### Definición
```go
type saver interface {
    Save() error
}

type getter interface {
    Get()
}

// Interface compuesta
type outputtable interface {
    saver
    Get()
}
```

### Implementación

Una interface se **implementa implícitamente** (sin `implements`):

```go
type Note struct {
    Title   string
    Content string
}

// Nota: no hay declaración "type Note implements saver"

func (n Note) Save() error {
    // implementación
    return nil
}

func (n Note) Get() {
    fmt.Println(n.Title, n.Content)
}
```

### Uso

```go
func saveData(data saver) error {
    return data.Save()
}

note := Note{Title: "My Note", Content: "..."}
err := saveData(note)  // funciona porque Note implementa saver

todo := Todo{Text: "My Task"}
err := saveData(todo)  // funciona si Todo también implementa saver
```

**Ventaja:** Un tipo puede satisfacer múltiples interfaces sin conocerlas de antemano.

### Interface Vacía

`interface{}` (o `any`) acepta **cualquier tipo**:

```go
func printAnything(value any) {
    fmt.Println(value)  // puede ser int, string, struct, etc.
}

printAnything(42)
printAnything("hello")
printAnything(3.14)
```

### Type Assertions (Afirmación de Tipo)

**Determinar el tipo exacto en tiempo de ejecución:**

```go
func process(value any) {
    // Type assertion: intenta afirmar que value es un string
    str, ok := value.(string)
    
    if ok {
        fmt.Println("Es un string:", str)
    } else {
        fmt.Println("No es un string")
    }
}

process("hello")  // Es un string: hello
process(42)       // No es un string
```

### Switch con Type

Determinar el tipo usando un `switch`:

```go
func process(value any) {
    switch v := value.(type) {
    case int:
        fmt.Println("Es un int:", v)
    case string:
        fmt.Println("Es un string:", v)
    case float64:
        fmt.Println("Es un float64:", v)
    default:
        fmt.Println("Tipo desconocido")
    }
}

process(42)        // Es un int: 42
process("hello")   // Es un string: hello
process(3.14)      // Es un float64: 3.14
```

**Nota:** En un `switch` con `.(type)`, solo puedes usar esta sintaxis. Fuera de un switch, debes usar type assertions.

---

### Ejemplo Completo: Múltiples Tipos

```go
// Sin generics (con interface{})
func addOld(a, b any) any {
    aInt, aIsInt := a.(int)
    bInt, bIsInt := b.(int)
    
    if aIsInt && bIsInt {
        return aInt + bInt
    }
    
    aStr, aIsStr := a.(string)
    bStr, bIsStr := b.(string)
    
    if aIsStr && bIsStr {
        return aStr + bStr
    }
    
    return nil
}

result := addOld(5, 10)      // 15 (int)
result := addOld("hi", "ya") // "hiya" (string)
```

---

## Generics

Los **Generics** permiten escribir funciones y tipos que funcionan con múltiples tipos sin perder type-safety.

### Sintaxis
```go
func add[T int | float64 | string](a, b T) T {
    return a + b
}

result := add(5, 10)          // 15 (int)
result := add(3.5, 2.5)       // 6.0 (float64)
result := add("hi", "ya")     // "hiya" (string)
```

**Restricción de tipos:** `T int | float64 | string` significa que `T` solo puede ser uno de esos tipos.

### Comparación: Generics vs Interface{}

**Con `interface{}`:**
```go
func add(a, b any) any {
    // necesita type assertion y manejo de errores
    aInt, aIsInt := a.(int)
    bInt, bIsInt := b.(int)
    if aIsInt && bIsInt {
        return aInt + bInt
    }
    // ... más type assertions
}
```

**Con Generics:**
```go
func add[T int | float64 | string](a, b T) T {
    return a + b  // más simple y type-safe
}
```

**Ventajas de Generics:**
- Type-safe en tiempo de compilación.
- Código más limpio sin type assertions.
- Mejor performance.

---

## Utilitarios Útiles

### Lectura de Entrada

**Lectura simple:**
```go
var choice int
fmt.Print("Enter choice: ")
fmt.Scan(&choice)
```

**Lectura de línea completa con `bufio`:**
```go
import (
    "bufio"
    "os"
    "strings"
)

reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')
text = strings.TrimSuffix(text, "\n")  // elimina el salto de línea
fmt.Println(text)
```

**Lectura múltiple con escáner:**
```go
scanner := bufio.NewScanner(os.Stdin)

fmt.Print("Ingresa tu nombre: ")
scanner.Scan()
name := scanner.Text()

fmt.Print("Ingresa tu edad: ")
scanner.Scan()
ageStr := scanner.Text()
age, _ := strconv.Atoi(ageStr)

fmt.Printf("Hola %s, tienes %d años\n", name, age)
```

**Diferencias:**
- **`fmt.Scan()`**: Simple pero requiere referencia (`&var`).
- **`bufio.NewReader().ReadString()`**: Lee hasta un carácter específico, retorna string.
- **`bufio.Scanner`**: Más flexible, parecido a otros lenguajes.

### Formato de Strings

**`fmt.Printf` (imprime):**
```go
fmt.Printf("Name: %v, Age: %d\n", "John", 30)
```

**`fmt.Sprintf` (devuelve string):**
```go
message := fmt.Sprintf("Hello %s", "World")
// message = "Hello World" (no se imprime)
```

### JSON

```go
import "encoding/json"

type Note struct {
    Title string `json:"title"`
    Content string `json:"content"`
}

// Serializar a JSON
data, err := json.Marshal(note)
os.WriteFile("note.json", data, 0644)

// Deserializar desde JSON
var note Note
json.Unmarshal(data, &note)
```

---

## Concurrencia (Goroutines)

Go tiene soporte nativo para concurrencia a través de **goroutines** y **channels**.

### Goroutines
Una goroutine es una función ejecutada de forma concurrente. Se inicia con la palabra clave `go`.

```go
func printMessages() {
    for i := 0; i < 5; i++ {
        fmt.Println("Message", i)
    }
}

func main() {
    go printMessages() // Se ejecuta de forma asíncrona en otra goroutine
    fmt.Println("Main execution")
    time.Sleep(1 * time.Second) // Evita que main termine antes que la goroutine
}
```

---

## REST API con Gin

Gin es uno de los frameworks web más populares de Go debido a su alto rendimiento.

### Servidor Básico y Rutas

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    server := gin.Default()

    // Endpoint GET
    server.GET("/events", getEvents)
    
    // Parámetros de ruta
    server.GET("/events/:id", getEvent)

    server.Run(":8080") // Iniciar el servidor en localhost:8080
}

func getEvents(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Lista de eventos"})
}
```

### Grupos de Rutas y Middlewares

Los middlewares permiten ejecutar código antes de llegar a los handlers.

```go
func Authenticate(ctx *gin.Context) {
    token := ctx.Request.Header.Get("Authorization")
    if token == "" {
        ctx.AbortWithStatusJSON(401, gin.H{"message": "No autorizado"})
        return
    }
    ctx.Next()
}

func RegisterRoutes(server *gin.Engine) {
    // Grupo protegido por middleware
    authenticated := server.Group("/")
    authenticated.Use(Authenticate)

    authenticated.POST("/events", createEvent)
    authenticated.PUT("/events/:id", updateEvent)
}
```

### Leer Datos (Bind JSON)

Para recibir JSON en peticiones POST/PUT:

```go
type Event struct {
    Name string `json:"name" binding:"required"`
}

func createEvent(ctx *gin.Context) {
    var event Event
    // BindJSON extrae el body http y lo mapea al Struct
    if err := ctx.BindJSON(&event); err != nil {
        ctx.JSON(400, gin.H{"message": "Datos inválidos"})
        return
    }
    // ...
}
```

---

## Bases de Datos con `database/sql`

Go soporta múltiples bases de datos con un paquete estándar `database/sql`, usado en combinación con un driver (ej: SQLite, PostgreSQL).

### Configuración y Conexión SQLite

```go
package db

import (
    "database/sql"
    _ "modernc.org/sqlite" // El underscore indica que importamos pero no usamos variables, solo inicializamos el driver
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "api.db")
    if err != nil {
        panic(err)
    }
}
```

### Queries y Executions

**`Exec`**: Para operaciones que no retornan filas (INSERT, UPDATE, DELETE, CREATE TABLE).

```go
query := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE
)`
_, err := DB.Exec(query)
```

**`Query` / `QueryRow`**: Para leer datos (SELECT).

```go
func GetEventByID(id int64) (*Event, error) {
    query := "SELECT id, name FROM events WHERE id = ?"
    row := DB.QueryRow(query, id)

    var event Event
    // Scan copia los valores devueltos en la fila a las variables
    err := row.Scan(&event.ID, &event.Name) 
    if err != nil {
        return nil, err
    }
    return &event, nil
}
```

---

## Recursos Útiles

- **`fmt` package:** Impresión y formato.
- **`errors` package:** Creación de errores.
- **`os` package:** Operaciones del sistema y archivos.
- **`strconv` package:** Conversión entre strings y números.
- **`time` package:** Manejo de fechas y horas.
- **`encoding/json` package:** Serialización JSON.
- **`bufio` package:** Lectura eficiente de entrada.