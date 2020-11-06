<template>
  <el-container>
    <el-header>
      <el-row>
        <el-col :span="20">
          <el-menu :default-active="activeIndex" cla ss="el-menu-demo" mode="horizontal"
            @select="handleSelect">
            <el-menu-item @click="$router.replace({name: 'Home'})">主页</el-menu-item>
            <el-menu-item @click="$router.replace({name: 'Predict'})">因子预测</el-menu-item>
            <el-menu-item @click="$router.replace({name: 'Create'})">因子创建</el-menu-item>
          </el-menu>
        </el-col>
        <el-col :span="4">
          <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal"
            @select="handleSelect">
            <el-menu-item @click="$router.replace({name: 'Login'})">登录</el-menu-item>
            <el-menu-item @click="$router.replace({name: 'Register'})">注册</el-menu-item>
          </el-menu>
        </el-col>
      </el-row>
    </el-header>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: '1',
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
  },
};
</script>

<style>
</style>
