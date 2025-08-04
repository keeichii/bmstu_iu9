# VBaumanke
## FORUM WEB APPLICATION
Project for the summer university internship at [Bauman Moscow State Technical
University](https://bmstu.ru/).

## Description
**VBaumanke** is a modern web-forum platform developed as part of the summer practice by students
of BMSTU. The project includes both frontend and backend components and is designed with a modern
architecture in mind.

### OVERVIEW

The application is a basic forum platform where users can register, log in,  
create posts, and leave comments. It is based on a PostgreSQL database, a  
FastAPI backend, and a static HTML/CSS/JS frontend.

### FEATURES

- User registration and JWT-based authentication  
- Post creation, editing, and deletion  
- Commenting system  
- Static frontend interface  
- Dockerized environment for easy deployment  

### TECHNOLOGIES

**Backend**  
- FastAPI  
- SQLAlchemy  
- JWT authentication  
- PostgreSQL  

**Frontend**  
- HTML5  
- CSS3  
- JavaScript  

**Others**  
- Docker  
- Docker Compose  


### PROJECT STRUCTURE
```
forum/
├── backend/ # Backend logic and routes (FastAPI)
├── frontend/ # Static files (HTML, CSS, JS)
├── init.sql # SQL initialization script
└── docker-compose.yml # Docker services configuration
```

### INSTALLATION

1. Ensure Docker and Docker Compose are installed.  
2. Clone the repository.  
3. Build and run the project using:  

```bash
   docker compose up --build -d
```

4. Visit the frontend in your browser:
	`http://localhost:5500`

5. The backend API is available at:
	`http://localhost:8000`

### Stop the Project
To stop the services, run:
```bash
	docker compose down
```

### CREDITS
Project prepared by:

1. [Khromov Matvey](https://gitflic.ru/user/khrmv)
2. [Yannaev Aleksandr](https://gitflic.ru/user/keeichi)