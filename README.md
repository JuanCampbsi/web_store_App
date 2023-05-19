<div style="width:100%; display: flex; gap: 15px; align-items: center;"> 
<h1> Web Store App </h1>
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="70" width="70" style="margin-bottom:1.5625rem"/>
</div>


### 🛠  Description   

</br>

The project aims to create a simple web application, using the Go programming language (Golang) without the use of frameworks. The application was developed with the aim of applying MVC (Model-View-Controller) concepts and performing basic CRUD operations (Create, Read, Update, Delete). PostgreSQL database is used for data storage.


## Preview 
User Interface.
</br>

<p align="center">
  <kbd>
 <img width="800" style="border-radius: 10px" height="400" src="" alt="Intro"> 
  </kbd>
  </br>
</p>

</br>

### ⌨ Database configuration
Create an .env file in the root of the project and define the database connection in it and put it in the db.go file. For example:

```bash
# .env
CONNECTION_DATABASE = host=your-host port=your-port user=your-user password=your-password dbname=database-name sslmode=require

# db/db.go
connection := os.Getenv("CONNECTION_DATABASE")

```

### ⌨ Installation
To use it, you need to clone the repository, install the dependencies and run the project.

```bash
# Open terminal/cmd and then Clone this repository
$ git clone https://github.com/JuanCampbsi/Web_Store_App.git

# Access project folder in terminal/cmd
$ cd Web_Store_App

# Install the dependencies
$ go mod tidy

# Run the application in development mode
$ go run main.go

```

</br>

### ⌨ Stack of technologies and libraries

-   [Golang](https://go.dev/doc/)
-   [PostgreSQL](https://www.postgresql.org/download/)

</br>

👨‍💻 **Author** 💻

Developed by [_Juan Campos_](https://www.linkedin.com/in/juancampos-ferreira/)