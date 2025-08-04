from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from database import get_db
from models import User
from schemas.user import UserCreate, UserOut
from utils.security import hash_password, verify_password
from fastapi.security import OAuth2PasswordRequestForm
from datetime import timedelta, datetime
from jose import jwt

router = APIRouter(prefix="/auth")

SECRET_KEY = "cbbf7f488d97e9418378c4d109a2c409a1c9b1f0a7fcd12867dfd2e20ac7a0c1"
ALGORITHM = "HS256"

@router.post("/register", response_model=UserOut)
def register(user: UserCreate, db: Session = Depends(get_db)):
    if db.query(User).filter(User.username == user.username).first():
        raise HTTPException(status_code=400, detail="Username taken")
    hashed_pw = hash_password(user.password)
    new_user = User(username=user.username, hashed_password=hashed_pw)
    db.add(new_user)
    db.commit()
    db.refresh(new_user)
    return new_user

@router.post("/login")
def login(
    form_data: OAuth2PasswordRequestForm = Depends(),
    db: Session = Depends(get_db)
):
    db_user = db.query(User).filter(User.username == form_data.username).first()
    if not db_user or not verify_password(form_data.password, db_user.hashed_password):
        raise HTTPException(status_code=401, detail="Invalid credentials")
    
    token_data = {
        "sub": db_user.username,
        "exp": datetime.utcnow() + timedelta(days=1)
    }
    token = jwt.encode(token_data, SECRET_KEY, algorithm=ALGORITHM)
    return {"access_token": token, "token_type": "bearer"}
