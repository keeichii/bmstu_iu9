# database.py
import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

# Берём строку подключения из переменной окружения DATABASE_URL
DATABASE_URL = os.getenv(
    "DATABASE_URL",
    "postgresql://forum_user:secret_pw@db:5432/forum"  # запасной вариант
)

# Создаём движок SQLAlchemy
engine = create_engine(DATABASE_URL)

# Фабрика сессий
SessionLocal = sessionmaker(
    autocommit=False,
    autoflush=False,
    bind=engine
)

# Базовый класс для декларативных моделей
Base = declarative_base()

# Зависимость для FastAPI
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
