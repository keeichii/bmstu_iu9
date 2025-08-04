from pydantic import BaseModel
from datetime import datetime

class CommentCreate(BaseModel):
    content: str

class CommentOut(BaseModel):
    id: int
    content: str
    timestamp: datetime
    author: str

    class Config:
        orm_mode = True