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
                <el-radio-button label="暗语"></el-radio-button>
                <el-radio-button label="密钥"></el-radio-button>
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
              <el-button type="primary" style="margin: 20px" @click="submitFetch"
                >取出密文</el-button
              >
            </div>
          </div>
        </div>
      </div>
    </div>
    <el-dialog title="完成" v-model="dialogTableVisible">
      <el-table :data="codenpin" style="width: 100%">
        <el-table-column label="密文已寄送" prop="type"> </el-table-column>
        <el-table-column label="" prop="name"> </el-table-column>
        <el-table-column fixed="right" label="" width="60">
          <template #default="scope">
            <el-button type="text" size="small" @click="handleClick(scope.row)"
              >复制到剪切板</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>


<style lang="scss" scoped>
@import "@/style/express.scss";
</style>

<script>
import { createExMessage } from "@/api/exMessage";
import { findExMessage } from "@/api/exMessage";

export default {
  data() {
    return {
      textarea1: "",
      textarea2: "",
      input1: "",
      input2: "",
      input3: "",
      select: "",
      radio1: "暗语",
      dialogTableVisible: false,
      codenpin: [
        {
          type: "密钥",
          name: "mymymymymymmymy",
        },
        {
          type: "暗语",
          name: "ayayayyayayayayay",
        },
      ],
    };
  },
  methods: {
    async submitMsg() {
      await createExMessage({ secret: this.textarea1 }).then((res) => {
      console.log(res)
      this.codenpin[0].name = res.data.pin
      this.codenpin[1].name = res.data.code
      });
      this.dialogTableVisible = true;
    },
    handleClick(val) {
      console.log(val.name);
      const input = document.createElement("input");
      document.body.appendChild(input);
      input.setAttribute("value", val.name);
      input.select();
      if (document.execCommand("copy")) {
        document.execCommand("copy");
        alert("复制成功");
      }
      document.body.removeChild(input);
    },
    async submitFetch() {
      await findExMessage({flag: this.radio1, query: this.textarea2 }).then((res) => {
      console.log(res)
      });
    }
  },
};
</script>
