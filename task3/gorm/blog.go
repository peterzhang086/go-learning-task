package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;not null;unique"`
	Email     string    `gorm:"size:100;not null;unique"`
	PostCount uint      `gorm:"size:100;comment:文章数量"` // 新增：文章数量统计字段
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// 一对多关系: 一个用户可以有多篇文章
	Posts []Post `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Post 文章模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Content       string    `gorm:"type:text;not null"`
	CommentStatus uint      `gorm:"size:100;comment:评论状态 1 为有评论 0为无评论"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`

	// 外键: 文章属于用户
	UserID uint `gorm:"not null"`
	// 多对一关系: 文章属于用户
	User User `gorm:"foreignKey:UserID"`
	// 一对多关系: 一篇文章可以有多个评论
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Comment 评论模型
type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// 外键: 评论属于文章
	PostID uint `gorm:"not null"`
	// 多对一关系: 评论属于文章
	Post Post `gorm:"foreignKey:PostID"`

	// 外键: 评论属于用户
	//UserID uint `gorm:"not null"`
	// 多对一关系: 评论属于用户
	//User User `gorm:"foreignKey:UserID"`
}

var gormDB *gorm.DB
var gorerr error

func init() {
	gormDB, gorerr = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if gorerr != nil {
		panic("failed to connect database")
	}
}

func main() {

	// 自动迁移 - 创建表
	gormDB.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 数据准备
	gormDB.Create(&User{Username: "张三", Email: "Jack@abc.com"})
	gormDB.Create(&User{Username: "李四", Email: "Rose@abc.com"})

	gormDB.Create(&Post{Content: "GORM 入门1", UserID: 1, CommentStatus: 1})
	gormDB.Create(&Post{Content: "GORM 入门2", UserID: 1, CommentStatus: 1})

	gormDB.Create(&Post{Content: "GORM 进阶1", UserID: 2, CommentStatus: 1})
	gormDB.Create(&Post{Content: "GORM 进阶2", UserID: 2, CommentStatus: 1})

	gormDB.Create(&Comment{Content: "GORM 入门1 很重要", PostID: 1})
	gormDB.Create(&Comment{Content: "GORM 入门1 入门很容易学", PostID: 1})

	gormDB.Create(&Comment{Content: "GORM 入门1 真的很容易么", PostID: 1})
	gormDB.Create(&Comment{Content: "GORM 入门2 如何入门", PostID: 2})

	gormDB.Create(&Comment{Content: "GORM进阶1 GORM对于什么应用那么重要", PostID: 3})
	gormDB.Create(&Comment{Content: "GORM进阶1 为什么重要", PostID: 3})

	gormDB.Create(&Comment{Content: "GORM进阶1 是不是很难啊", PostID: 3})
	gormDB.Create(&Comment{Content: "GORM进阶2 这么难怎么学", PostID: 3})

	gormDB.Create(&Comment{Content: "GORM进阶2 终极要义", PostID: 4})

	// 关联查询
	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var posts []Post
	gormDB.Debug().Preload("Comments").Where("user_id = ?", 1).Find(&posts)
	jsonByte, _ := json.Marshal(posts)
	fmt.Println(string(jsonByte))

	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var post Post
	gormDB.Debug().Model(&Comment{}).Select("*").Group("post_id").Order("count(post_id) desc").Limit(1).Find(&post)
	jsonByte, _ = json.Marshal(post)
	fmt.Println(string(jsonByte))

	// 新增文章，自动更新用户文章数量
	gormDB.Create(&Post{Content: "GORM 入门3", UserID: 1, CommentStatus: 1})

	//删除评论，自动更新文章评论状态
	var comment Comment
	gormDB.Debug().Where("id = ?", 8).Find(&comment)
	gormDB.Debug().Where("id = ?", comment.ID).Delete(&comment)
}

// AfterCreate - Post 创建后的钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户文章数量
	// 使用 gorm.Expr 来原子性地增加计数，避免并发问题
	result := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1"))
	if result.Error != nil {
		return result.Error
	}
	fmt.Printf("用户 ID %d 的文章数量已更新，当前文章数: +1\n", p.UserID)
	return nil
}

// AfterDelete - Comment 删除后得钩子函数
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 更新文章评论状态  如果没有评论了，则设置为0 无评论状态
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", 0)
		fmt.Println("删除评论后，更新文章评论状态为无评论")
	} else {
		fmt.Println("删除评论后，文章仍有评论", count)
	}
	return
}
