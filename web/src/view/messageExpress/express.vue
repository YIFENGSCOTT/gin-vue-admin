<template>
  <div>
    <div class="container">
      <!-- Pricing column -->
      <div class="container__column" id="send__column">
        <div class="card" id="send__card">
          <!-- Cover -->
          <div class="card__title">寄送一条秘文</div>

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
            >生成密钥和暗语</el-button
          >
        </div>
      </div>

      <div class="container__column">
        <div class="card" id="retrieve__card">
          <!-- Cover -->
          <div class="card__title">取回一条秘文</div>
          <div class="selections">
            <div style="margin: 20px">
              <el-radio-group v-model="radio1">
                <el-radio-button label="密钥"></el-radio-button>
                <el-radio-button label="暗语"></el-radio-button>
              </el-radio-group>
            </div>
            <div class="card__content">
              <el-input
                type="textarea"
                autosize
                placeholder="请输入"
                v-model="textarea2"
                class="send__area"
                style="margin-top: -10px"
              >
              </el-input>
            </div>
            <div>
              <el-button type="primary" style="margin: 20px"
                >取出密文</el-button
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<style lang="scss" scoped>
@import "@/style/express.scss";
</style>

<script>
import { createExMessage } from "@/api/exMessage";

export default {
  data() {
    return {
      textarea1: "",
      textarea2: "",
      input1: "",
      input2: "",
      input3: "",
      select: "",
      radio1: "New York",
    };
  },
  methods: {
    async submitMsg() {
      await createExMessage({secret: this.textarea1}).then((res) => {
        alert(JSON.stringify(res));
      });
    },
  },
};
</script>
