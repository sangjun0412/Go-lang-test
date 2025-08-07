package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-contrib/cors"
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
	Port             int    `json:"port"`          // 포트번호
	ReceivedTime     string `json:"time"`          // 시간
	TotalCount       int    `json:"total_count"`   // 총 합계
	ErrorCount       int    `json:"error_count"`   // 오류
	CurrentCount     int    `json:"current_count"` // 현재 수신값
	ConnectionStatus string `json:"status"`        // 연결상태
}

// 사용자 정보를 메모리에 저장할 슬라이스
var users []User
var nextID int = 1 // 사용자 ID 자동 증가 용도

// FepData를 저장할 맵 (포트번호를 키로 사용)
var fepDataMap = make(map[int]FepData)

func main() {
	// 기본 라우터 생성
	router := gin.Default()

	// 기본 CORS 설정으로 모든 origin 허용
	router.Use(cors.Default())

	// 사용자 목록 조회 API
	router.GET("/users", getUsers)

	// 사용자 등록 API
	router.POST("/users", createUser)

	// 데이터 조회 API
	router.GET("/data", getAllFepdata)

	// 데이터 등록 API
	router.POST("/receive", receiveFepdata)

	// 8080 포트에서 서버 실행
	router.Run(":8080")
}

// 사용자 목록 조회 핸들러
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users) // 모든 사용자 정보를 JSON으로 반환
}

// 모든 FepData를 슬라이스 형태로 반환
func getAllFepdata(c *gin.Context) {
	values := make([]FepData, 0, len(fepDataMap))
	for _, data := range fepDataMap {
		values = append(values, data)
	}
	c.JSON(http.StatusOK, values)
}

func handleData(c *gin.Context) {
	//  C 프로그램 실행 (동기 실행)
	//  ex: your_c_program이 exit code 0으로 종료되면, 정상 수행
	cmd := exec.Command("./test") // 파일명과 필요 시 인자 추가
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": string(out)})
		return
	}

	//  실행 성공 시 응답 데이터로 C 프로그램 출력 반환
	c.JSON(http.StatusOK, gin.H{
		"message": "C program executed successfully",
		"output":  string(out),
	})
}

func handleDataAsync(c *gin.Context) {
	go func() {
		out, err := exec.Command("./your_c_program").CombinedOutput()
		// 필요시 로깅하거나 결과 저장 처리
		if err != nil {
			log.Println("C program error:", err, string(out))
		} else {
			log.Println("C output:", string(out))
		}
	}()
	c.JSON(http.StatusAccepted, gin.H{"message": "C program started"})
}

// 사용자 등록 핸들러
func receiveFepdata(c *gin.Context) {
	var newData FepData

	// 요청 body에서 JSON 데이터를 파싱하여 newUser에 바인딩
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}

	// 포트번호를 키로 하여 데이터 저장 (기존 키면 업데이트, 아니면 추가됨)
	fepDataMap[newData.Port] = newData
	c.JSON(http.StatusCreated, newData)
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
