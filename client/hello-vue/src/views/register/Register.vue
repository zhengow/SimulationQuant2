<template>
  <div class="register">
    <el-row>
      <el-col :span="8" :offset="8">
        <el-card class="box-card">
          <el-form class="login-container">
            <h1 class="title" align="center">用户注册</h1>
            <el-form-item>
              <el-link type="primary" :underline="false" @click="$router.replace({name: 'Login'})">
                已有账号?</el-link>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input type="text" v-model="user.Name" placeholder="Enter username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="user.Pwd" placeholder="Enter password"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" style="width: 100%;" @click="register">登陆</el-button>
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
    register() {
      const api = 'http://49.234.229.140:8082/api/register';
      this.axios.post(api, this.user).then((res) => {
        console.log(res);
      });
      console.log('register');
    },
  },
};
</script>
