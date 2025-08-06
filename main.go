package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 사용자 정보를 담을 구조체 정의
type User struct {
	ID    int    `json:"id"`    // 사용자 ID
	Name  string `json:"name"`  // 이름
	Email string `json:"email"` // 이메일
}

// 데이터 정보를 담을 구조체 정의
type FepData struct {
	Port             int    `json:"port"`           // 포트번호
	ReceivedTime     string `json:"time"`           // 시간
	TotalCount       int    `json:"total_count"`    // 총 합계
	ErrorCount       int    `json:"error_count"`    // 오류
	CurrentCount     int    `json: "current_count"` // 현재 수신값
	ConnectionStatus string `json: "status"`        // 연결상태
}

// 사용자 정보를 메모리에 저장할 슬라이스
var users []User
var nextID int = 1 // 사용자 ID 자동 증가 용도

var fepDatas []FepData

func main() {
	// 기본 라우터 생성
	router := gin.Default()

	// 사용자 목록 조회 API
	router.GET("/users", getUsers)

	// 사용자 등록 API
	router.POST("/users", createUser)

	// 데이터 조회 API
	router.GET("/receive", getFepData)

	// 데이터 등록 API
	router.POST("/receive", receiveFepdata)

	// 8080 포트에서 서버 실행
	router.Run(":8080")
}

// 사용자 목록 조회 핸들러
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users) // 모든 사용자 정보를 JSON으로 반환
}

// 사용자 목록 조회 핸들러
func getFepData(c *gin.Context) {
	c.JSON(http.StatusOK, fepDatas) // 모든 사용자 정보를 JSON으로 반환
}

// 사용자 등록 핸들러
func receiveFepdata(c *gin.Context) {
	var newFepData FepData

	// 요청 body에서 JSON 데이터를 파싱하여 newUser에 바인딩
	if err := c.BindJSON(&newFepData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}

	// 메모리에 사용자 추가
	fepDatas = append(fepDatas, newFepData)

	// 등록된 사용자 정보를 반환
	c.JSON(http.StatusCreated, newFepData)
}

// 사용자 등록 핸들러
func createUser(c *gin.Context) {
	var newUser User

	// 요청 body에서 JSON 데이터를 파싱하여 newUser에 바인딩
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}

	// ID 자동 할당
	newUser.ID = nextID
	nextID++

	// 메모리에 사용자 추가
	users = append(users, newUser)

	// 등록된 사용자 정보를 반환
	c.JSON(http.StatusCreated, newUser)
}
