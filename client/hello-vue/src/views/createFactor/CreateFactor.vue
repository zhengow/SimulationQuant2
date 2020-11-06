<template>
  <div class="login">
    <el-row>
      <el-col :span="8" :offset="8">
        <el-card class="box-card">
          <el-form class="login-container">
            <h1 class="title" align="center">用户登录</h1>
            <el-form-item label="用户名">
              <el-input type="text" v-model="user.Name" placeholder="Enter username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="user.Pwd" placeholder="Enter password"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" style="width: 100%;" @click="login()">登陆</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        Name: '',
        Pwd: '',
      },
    };
  },
  methods: {
    login() {
      const api = 'http://49.234.229.140:8082/api/login';
      this.axios.post(api, this.user).then((res) => {
        console.log(res);
        if (res.data.Data === true) {
          // 跳转页面
        }
      });
    },
  },
};
</script>
