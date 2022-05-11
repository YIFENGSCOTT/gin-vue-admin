<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="信息内容:">
          <el-input v-model="formData.threadContent" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="父信息流:">
          <el-input v-model.number="formData.fatherThread" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密钥Id:">
          <el-input v-model.number="formData.encryptionKeyId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="浏览次数:">
          <el-input v-model.number="formData.visited" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备用:">
          <el-input v-model="formData.beiyong" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Thread'
}
</script>

<script setup>
import {
  createThread,
  updateThread,
  findThread
} from '@/api/thread'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        threadContent: '',
        fatherThread: 0,
        encryptionKeyId: 0,
        visited: 0,
        beiyong: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findThread({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rethread
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createThread(formData.value)
          break
        case 'update':
          res = await updateThread(formData.value)
          break
        default:
          res = await createThread(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
