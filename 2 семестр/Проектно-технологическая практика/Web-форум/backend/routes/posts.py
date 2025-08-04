from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from schemas.post import PostCreate, PostOut
from models import Post, User
from database import get_db
from .get_current_user import get_current_user

router = APIRouter(
    prefix="/posts",
    tags=["posts"]
)

@router.post("/", response_model=PostOut)
def create_post(
    post: PostCreate,
    db: Session = Depends(get_db),
    current_user: User = Depends(get_current_user)
):
    new_post = Post(title=post.title, content=post.content, author=current_user)
    db.add(new_post)
    db.commit()
    db.refresh(new_post)
    return PostOut(
        id=new_post.id,
        title=new_post.title,
        content=new_post.content,
        timestamp=new_post.timestamp,
        author=new_post.author.username,
        comments=[]
    )

@router.get("/", response_model=list[PostOut])
def get_posts(db: Session = Depends(get_db)):
    posts = db.query(Post).order_by(Post.id.desc()).all()
    return [
        PostOut(
            id=p.id,
            title=p.title,
            content=p.content,
            timestamp=p.timestamp,
            author=p.author.username,
            comments=[]
        )
        for p in posts
    ]
