<template>
  <div class="h-screen bg-green-900 flex flex-col items-center justify-between text-white p-4">
    <!-- 顶部状态栏 -->
    <div class="w-full flex justify-between items-center text-sm">
      <div class="flex items-center gap-2">
        <div class="w-10 h-10 bg-gray-500 rounded-full flex items-center justify-center">?</div>
        <span>房间: {{ roomId }}</span>
      </div>
      <div class="text-xs">
        <span>电池 80%</span> | <span>21:00</span>
      </div>
    </div>

    <!-- 牌桌区域 -->
    <div class="relative w-full h-3/5 border-2 border-green-800 rounded-full flex items-center justify-center">
      <!-- 桌面中间骰子/庄家 -->
      <div class="w-20 h-20 bg-green-700/50 rounded-lg flex items-center justify-center">
        <span class="text-2xl font-bold">东</span>
      </div>
    </div>

    <!-- 玩家手牌区域 (移动端) -->
    <div class="w-full flex flex-col items-center gap-2">
      <div class="flex gap-1 overflow-x-auto p-2">
        <div v-for="card in hand" :key="card"
             class="w-10 h-14 bg-white text-black rounded-sm border-2 border-gray-300 flex items-center justify-center font-bold text-xl active:translate-y-[-10px] transition-transform">
          {{ formatCard(card) }}
        </div>
      </div>
      <!-- 操作按钮 -->
      <div class="flex gap-4">
        <button v-if="canHu" class="px-6 py-2 bg-red-600 rounded-full font-bold shadow-lg">胡</button>
        <button v-if="canGang" class="px-6 py-2 bg-yellow-600 rounded-full font-bold shadow-lg">杠</button>
        <button v-if="canPeng" class="px-6 py-2 bg-blue-600 rounded-full font-bold shadow-lg">碰</button>
        <button class="px-6 py-2 bg-gray-600 rounded-full font-bold shadow-lg">过</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const roomId = ref('8888')
const hand = ref([11, 11, 11, 22, 23, 24, 31, 32, 33, 35, 36, 37, 39])
const canHu = ref(false)
const canGang = ref(false)
const canPeng = ref(true)

const formatCard = (c) => {
  const t = Math.floor(c / 10)
  const v = c % 10
  const names = ['', '万', '条', '饼']
  return v + names[t]
}
</script>

<style scoped>
/* 微信浏览器隐藏滚动条 */
::-webkit-scrollbar {
  display: none;
}
</style>
