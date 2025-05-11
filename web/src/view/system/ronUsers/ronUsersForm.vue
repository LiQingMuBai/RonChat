
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="deletedAt字段:" prop="deletedAt">
    <el-date-picker v-model="formData.deletedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="用户登录名:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入用户登录名" />
</el-form-item>
        <el-form-item label="用户登录密码:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入用户登录密码" />
</el-form-item>
        <el-form-item label="男女 1男 0女:" prop="sex">
    <el-input v-model.number="formData.sex" :clearable="true" placeholder="请输入男女 1男 0女" />
</el-form-item>
        <el-form-item label="用户是否被冻结 1正常 2冻结:" prop="enable">
    <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入用户是否被冻结 1正常 2冻结" />
</el-form-item>
        <el-form-item label="用户头像:" prop="headerImg">
    <el-input v-model="formData.headerImg" :clearable="true" placeholder="请输入用户头像" />
</el-form-item>
        <el-form-item label="用户手机号:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入用户手机号" />
</el-form-item>
        <el-form-item label="用户邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入用户邮箱" />
</el-form-item>
        <el-form-item label="用户telegram:" prop="telegram">
    <el-input v-model="formData.telegram" :clearable="true" placeholder="请输入用户telegram" />
</el-form-item>
        <el-form-item label="roomID:" prop="roomId">
    <el-input v-model="formData.roomId" :clearable="true" placeholder="请输入roomID" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createRonUsers,
  updateRonUsers,
  findRonUsers
} from '@/api/system/ronUsers'

defineOptions({
    name: 'RonUsersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
            username: '',
            password: '',
            sex: undefined,
            enable: undefined,
            headerImg: '',
            phone: '',
            email: '',
            telegram: '',
            roomId: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRonUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createRonUsers(formData.value)
               break
             case 'update':
               res = await updateRonUsers(formData.value)
               break
             default:
               res = await createRonUsers(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
