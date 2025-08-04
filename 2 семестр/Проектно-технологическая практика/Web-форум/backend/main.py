from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from routes import auth, posts, comments
from utils.init_db import init_db

app = FastAPI(
    title="Forum API",
    description="Backend for forum: auth, posts, comments",
    version="1.0.0"
)

# Разрешаем CORS для фронтенда
app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5500"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Подключаем роутеры (префиксы внутри файлов)
app.include_router(auth.router)
app.include_router(posts.router)
app.include_router(comments.router)

# Создание таблиц при запуске
@app.on_event("startup")
def on_startup():
    init_db()
