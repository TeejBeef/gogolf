<script lang="ts" setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { GetTeeTimes } from '../../wailsjs/go/main/App'
import { time } from '../../wailsjs/go/models'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

const teetime = reactive({
  courses: ["", ""],
  date: Date,
  time: time,
  results: "",
})

function getteetimes() {
  GetTeeTimes(teetime.courses, teetime.time).then(result => {
    console.log(result)
  })
}

</script>

<template>
  <main>
    <!-- <div id="result" class="result">{{ data.resultText }}</div> -->
    <div id="input" class="input-box">
      <input id="teetimescourses" v-model="teetime.courses" autocomplete="off" class="input" type="options" />
      <input id="teetimesdate" v-model="teetime.date" autocomplete="off" class="input" type="date" />
      <input id="teetimestime" v-model="teetime.time" autocomplete="off" class="input" type="time" />
      <button class="btn" @click="getteetimes">GetTeeTimes</button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: grab;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
