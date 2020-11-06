<template>
  <el-container>
    <el-header>
      <el-row>
        <el-col :span="6" :offset="1">
          <el-date-picker v-model="date" type="month" placeholder="选择月份查看当月预测">
          </el-date-picker>
        </el-col>
        <el-col :span="6">
          <el-select v-model="factor" placeholder="请选择因子">
            <el-option
              v-for="item in factors"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="6">
          <el-button type="primary" icon="el-icon-search" @click="test">搜索</el-button>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="date" label="日期" width="180">
        </el-table-column>
        <el-table-column prop="stock" label="股票" width="180">
        </el-table-column>
        <el-table-column prop="return" label="收益率">
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      date: null,
      tableData: [],
      factors: [{
        value: '1',
        label: '一号因子',
      }, {
        value: '2',
        label: '二号因子',
      }, {
        value: '3',
        label: '三号因子',
      }, {
        value: '4',
        label: '四号因子',
      }, {
        value: '5',
        label: '五号因子',
      }],
      factor: '',
    };
  },
  created() {
    this.test1();
  },
  methods: {
    test1() {
      console.log('query data');
    },
    test() {
      console.log(this.factor);
      this.tableData = [];
      const year = this.date.getYear() + 1900;
      const month = this.date.getMonth() + 1;
      const params = { date: `${year}-${month}` };
      console.log(params);
      const api = 'http://49.234.229.140:8082/api/predict';
      this.axios.post(api, params).then((res) => {
        res.data.Data.forEach((item) => {
          this.tableData.push(item);
        });
        console.log(res.data.Data);
      });
    },
  },
};
</script>

<style>
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}
.el-dropdown-link {
  cursor: pointer;
  color: #ffffff;
}
.el-button {
  background-color: white;
  color: #409eff;
  border-color: rgb(199, 199, 199);
}
.el-icon-arrow-down {
  font-size: 12px;
}
.el-date-picker {
  border-color: #409eff;
}
</style>
