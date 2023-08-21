package reqDto

/**
 * @ClassName authDto
 * @Description TODO
 * @Author khr
 * @Date 2023/8/1 14:15
 * @Version 1.0
 */
type Login struct {
	Email    string `json:"email" `
	Phone    string `json:"phone" `
	Password string `json:"password" binding:"required"`
	Method   string `json:"method" binding:"required" gorm:"one of email,phone"`
}
