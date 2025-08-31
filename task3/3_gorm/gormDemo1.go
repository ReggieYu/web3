package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetUserPostsWithComments(db *gorm.DB, userId int) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

func (p *Post) AfterCreate(db *gorm.DB) error {
	result := db.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", gorm.Expr("post_count + ?", 1))
	return result.Error
}

func (c *Comment) AfterDelete(db *gorm.DB) error {
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
	if count == 0 {
		db.Model(&Post{}).Where("id = ?", c.PostId).UpdateColumn("comment_status", "无评论")
	}

	return nil
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection failed")
	}

	//题目1：模型定义
	//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	//要求 ：
	//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	//编写Go代码，使用Gorm创建这些模型对应的数据库表。
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("database migrate failed")
	}

	//题目2：关联查询
	//基于上述博客系统的模型定义。
	//要求 ：
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	posts, err := GetUserPostsWithComments(db, 1)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		fmt.Printf("post title: %s \n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("comment content: %s\n", comment.Content)
		}
	}

	//题目3：钩子函数
	//继续使用博客系统的模型。
	//要求 ：
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	fmt.Println("\n1. post count before create new post：")
	var userBef User
	db.Preload("Posts").Where("id = ?", 1).First(&userBef)
	fmt.Printf("user： %s \npost count: %d", userBef.Username, userBef.PostCount)

	// create new post
	newPost := Post{
		Title:   "test post",
		Content: "test content",
		UserId:  1,
	}

	err = db.Create(&newPost).Error
	if err != nil {
		fmt.Printf("\ncreate new post failed: %v\n", err)
	} else {
		fmt.Println("\ncreate new post success")
	}

	fmt.Println("\n2. post count after crate new post: ")
	var userAf User
	db.Preload("Posts").Where("id = ?", 1).First(&userAf)
	fmt.Printf("user: %s\npost count: %d", userAf.Username, userAf.PostCount)

	// delete comment
	fmt.Println("\n3. post status before delete comment")
	var postBef Post
	db.Preload("Comments").First(&postBef, 1)
	fmt.Printf("post title: %s comment status: %s comment count: %d", postBef.Title, postBef.CommentStatus, len(postBef.Comments))
	newComments := Comment{
		Content: "test comment content",
		PostId:  1,
	}

	db.Create(&newComments)
	fmt.Println("\ncreate a new comment")
	var postAf Post
	db.Preload("Comments").First(&postAf, 1)
	fmt.Printf("post title: %s comment status: %s comment count: %d", postAf.Title, postAf.CommentStatus, len(postAf.Comments))

	if err := db.Delete(&newComments).Error; err != nil {
		fmt.Printf("delete comment failed %v\n", err)
	} else {
		fmt.Println("\n4. post status after delete comment:")
	}

	var postDel Post
	db.Preload("Comments").First(&postDel, 1)
	fmt.Printf("post title: %s comment status: %s comment count: %d", postDel.Title, postDel.CommentStatus, len(postDel.Comments))
}
