#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <time.h>
#include <curl/curl.h>

int main(void)
{
    CURL *curl;
    CURLcode res;
    int port, total=1500, current;
    char buffer[512];
    const char *json_data;
    time_t now = time(NULL);
    struct tm *t = localtime(&now);
    char timebuf[64];
    strftime(timebuf, sizeof(timebuf), "%Y-%m-%dT%H:%M:%S", t);

/*
    const char *json_data = "{\"id\":1,\"name\":\"상준222\",\"email\":\"sangjun@example.com\"}";
*/
/*
    const char *json_data =
        "{"
        "\"port\": 8084,"
        "\"time\": \"2025-08-06T13:00:00\","
        "\"total_count\": 1432123,"
        "\"error_count\": 25,"
        "\"current_count\": 150,"
        "\"status\": \"connected\""
        "}";
*/
    struct curl_slist *headers = NULL;
    
  
    curl = curl_easy_init();
    if (curl) {
        /* 헤더 설정 */
        headers = curl_slist_append(headers, "Content-Type: application/json");
        headers = curl_slist_append(headers, "Accept: application/json");

        /* 요청할 URL 설정 */
        curl_easy_setopt(curl, CURLOPT_URL, "http://localhost:8080/receive");

        /* 강제 IPv4로 전송*/
        curl_easy_setopt(curl, CURLOPT_IPRESOLVE, CURL_IPRESOLVE_V4);

        /* POST 방식으로 설정 */
        curl_easy_setopt(curl, CURLOPT_POST, 1L);

        for (port = 11001; port <= 11023; ++port) {
            current = 10;
            total += current;
    
            // 충분히 큰 버퍼 할당
            snprintf(buffer, sizeof(buffer),
                "{"
                "\"port\": %d,"
                "\"time\": \"%s\","
                "\"total_count\": %d,"
                "\"error_count\": %d,"
                "\"current_count\": %d,"
                "\"status\": \"%s\""
                "}",
                port,
                timebuf,
                total,
                25,
                current,
                "connected"
            );
            json_data = buffer;
printf("%s\n", json_data);
            /* JSON 데이터 전송 */
            curl_easy_setopt(curl, CURLOPT_POSTFIELDS, json_data);

            /* 헤더 적용 */
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

        /* 응답을 stdout에 출력 */
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, NULL);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, stdout);

        /* 요청 수행 */
        res = curl_easy_perform(curl);

        /* 결과 확인 */
        if (res != CURLE_OK)
            fprintf(stderr, "curl_easy_perform() 실패: %s\n", curl_easy_strerror(res));
        }
        /* 정리 */
        curl_slist_free_all(headers);
        curl_easy_cleanup(curl);
    }

    return 0;
}
