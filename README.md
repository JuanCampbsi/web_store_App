<div style="width:100%; display: flex; align-items: center;">
  <h1>Web Store App
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="60" width="70" style="margin-bottom: -15px; z-index: -10; margin-left: 1.25rem"/>
  </h1> 
</div>



### 🛠  Description   

</br>

The project aims to create a simple web application using the Go programming language (Golang) without the use of frameworks. The application follows the MVC (Model-View-Controller) architecture and performs basic CRUD operations (Create, Read, Update, Delete). PostgreSQL database is used for data storage.

The project also includes the integration of FlatBuffers, a serialization library, for efficient data storage and retrieval. FlatBuffers provides a compact binary format that allows for fast serialization and deserialization of data, making it suitable for applications that require high performance and low memory usage.

By leveraging FlatBuffers, the application can optimize data serialization and minimize memory overhead, improving overall efficiency and scalability.


## Preview 
User Interface.
</br>

<p align="center">
  <kbd>
 <img width="800" style="border-radius: 10px" height="400" src="https://github.com/JuanCampbsi/Web_Store_App/blob/77a7f78d7dffbec5d261443a19d88322a80d5052/templates/assets/preview.gif" alt="Intro"> 
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

# To change the schema files and generate a new buff, install FlatBuffers on your machine and run the command below
$ flatc --go schemas/products.fbs

```

</br>	

### ⌨ Stack of technologies and libraries

-   [Golang](https://go.dev/doc/) - version 1.20
-   [PostgreSQL](https://www.postgresql.org/download/) - version 10.22 
-   [FlatBuffers](https://flatbuffers.dev/) - version 23.5.9 
</br>

👨‍💻 **Author** 💻

Developed by [_Juan Campos_](https://www.linkedin.com/in/juancampos-ferreira/)