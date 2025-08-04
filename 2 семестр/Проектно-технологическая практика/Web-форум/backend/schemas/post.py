from pydantic import BaseModel
from datetime import datetime
from typing import List
from schemas.comment import CommentOut

class PostCreate(BaseModel):
    title: str
    content: str

class PostOut(BaseModel):
    id: int
    title: str
    content: str
    timestamp: datetime
    author: str
    comments: List[CommentOut] = []

    class Config:
        orm_mode = True
