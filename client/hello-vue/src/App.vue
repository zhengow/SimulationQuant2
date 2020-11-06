<template>
  <div id="app">
    <navbar />
    <el-main>
      <router-view />
    </el-main>
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue';

export default {
  components: { Navbar },
  data() {
    return {};
  },
};
</script>
