<template>
  <div>
    <!-- 조회 버튼 누를 때마다 fetchData 호출 -->
    <button @click="fetchData">조회</button>
    <h2>FEP 포트 모니터링</h2>
    <!-- 데이터 배열 반복 렌더링 -->
    <table>
      <thead>
        <tr><th>포트번호</th><th>수신시간</th><th>총합</th><th>에러카운트</th><th>현재수신된값</th><th>Status</th></tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.port">
          <td>{{ item.port }}</td>
          <td>{{ item.time }}</td>
          <td>{{ item.total_count }}</td>
          <td>{{ item.error_count }}</td>
          <td>{{ item.current_count }}</td>
          <td>{{ item.status }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
// Composition API 방식
import { ref, onMounted } from 'vue';
import axios from 'axios';

const items = ref([]); // API 응답 데이터를 저장할 변수

async function fetchData() {
  try {
    const res = await axios.get('http://localhost:8080/data');
    console.log('fetchData response:', res);        // 전체 응답 객체 출력
    console.log('response.data:', res.data);        // 실제 JSON payload 출력
    items.value = res.data; // res.data는 FepData 배열S

  } catch (err) {
    console.error('API 호출 중 오류 발생:', err);
  }
}

onMounted(fetchData); // 컴포넌트가 마운트될 때 데이터 요청
</script>

<style>
table { width: 100%; border-collapse: collapse; }
th, td { border: 1px solid #aaa; padding: 6px; text-align: center; }
button { margin-bottom: 10px; }
</style>
