from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from database import get_db
from models import Comment, User, Post
from schemas.comment import CommentCreate, CommentOut
from typing import List
from .get_current_user import get_current_user

router = APIRouter(prefix="/comments")

@router.post("/{post_id}", response_model=CommentOut)
def add_comment(
    post_id: int,
    comment: CommentCreate,
    db: Session = Depends(get_db),
    current_user: User = Depends(get_current_user)
):
    post = db.query(Post).filter(Post.id == post_id).first()
    if not post:
        raise HTTPException(status_code=404, detail="Post not found")
    new_comment = Comment(content=comment.content, author=current_user, post=post)
    db.add(new_comment)
    db.commit()
    db.refresh(new_comment)
    return CommentOut(
        id=new_comment.id,
        content=new_comment.content,
        timestamp=new_comment.timestamp,
        author=new_comment.author.username
    )

@router.get("/{post_id}", response_model=List[CommentOut])
def get_comments(post_id: int, db: Session = Depends(get_db)):
    comments = db.query(Comment).filter(Comment.post_id == post_id).all()
    return [
        CommentOut(
            id=c.id,
            content=c.content,
            timestamp=c.timestamp,
            author=c.author.username
        ) for c in comments
    ]
