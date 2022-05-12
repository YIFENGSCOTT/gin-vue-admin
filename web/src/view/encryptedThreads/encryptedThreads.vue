<template>
  <div>
    <div>
      <el-row>
        <el-col
          v-for="(o, index) in 1"
          :key="o"
          :span="8"
          :offset="index > 0 ? 2 : 0"
        >
          <el-card
            :body-style="{ padding: '0px' }"
            style="height: 60px; width: 230px"
          >
            <img
              src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              class="image"
              style="display: inline-block; vertical-align: middle"
            />
            <div style="padding: 14px; display: inline-block">
              <span
                style="
                  max-width: 5em;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  white-space: nowrap;
                "
                >密钥已装载</span
              >
              <div style="display: inline-block; margin-left: 40px">
                <el-button
                  type="text"
                  class="button"
                  @click="showChangeKeyDialog"
                  >更换</el-button
                >
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div class="container">
      <!-- Pricing column -->
      <div class="container__column" id="send__column">
        <div class="card" id="send__card">
          <!-- Cover -->
          <div class="card__title">创建信息流</div>

          <!-- Content -->
          <div class="card__content">
            <el-input
              type="textarea"
              autosize
              placeholder="Please input"
              v-model="textarea1"
              class="send__area"
            >
            </el-input>
          </div>
          <el-button type="primary" style="margin: 20px" @click="submitMsg"
            >发布</el-button
          >
        </div>
      </div>
    </div>
    <el-dialog title="更换密钥" v-model="dialogTableVisible">
      <div style="margin-top: 20px">
        <el-input v-model="input1" placeholder="Please input">
          <template #prepend>密钥</template>
        </el-input>
      </div>
      <div style="display: flex; justify-content: flex-end; margin-top: 20px">
        <el-button type="" style="margin: 20px" @click="newKey"
          >用该账号身份新建密钥</el-button
        >

        <el-button
          type="primary"
          style="margin: 20px"
          @click="showChangeKeyDialog"
          >更换密钥</el-button
        >
      </div>
    </el-dialog>

    <el-dialog title="新密钥" v-model="dialogTable2Visible">
      <div>
        为了保证通讯安全，请输入一个用来接收密钥的邮箱地址
      </div>
      <div style="margin-top: 20px">
        <el-input v-model="email" placeholder="Please input">
          <template #prepend>邮箱:</template>
        </el-input>
      </div>

      <el-button
        type="primary"
        @click="submitNewKey" style="margin-top: 20px"
        >创建密钥</el-button
      >
    </el-dialog>
  </div>
</template>


<style lang="scss" scoped>
@import "@/style/threads.scss";
</style>

<script>
import { createExMessage } from "@/api/exMessage";
import { findExMessageByCode } from "@/api/exMessage";
import { findExMessageByPin } from "@/api/exMessage";
import { ref } from "vue";
import { createEncryptionKey } from "@/api/encryptionKey";

const currentDate = ref(new Date());

export default {
  data() {
    return {
      textarea1: "",
      textarea2: "",
      input1: "",
      input2: "",
      input3: "",
      email: "",
      newKeyName:"",
      select: "",
      radio1: "暗语",
      dialogTableVisible: false,
      dialogTable2Visible: false,
      codenpin: [
        {
          type: "密钥",
          name: "",
        },
        {
          type: "暗语",
          name: "",
        },
      ],
      secretnvisits: [
        {
          type: "密文",
          name: "",
        },
        {
          type: "取出次数",
          name: "",
        },
      ],
    };
  },
  methods: {
    showChangeKeyDialog() {
      this.dialogTableVisible = true;
    },
    newKey() {
      this.dialogTable2Visible = true;
    },
    async submitNewKey(){
      await createEncryptionKey({ Beiyong: this.email }).then((res) => {
        console.log(res);
      });
    }
  },
};
</script>


