from database import SessionLocal
from models import Post

db = SessionLocal()
db.query(Post).delete()
db.commit()
print("✅ Все посты удалены.")
