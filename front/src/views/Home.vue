<template>
  <div>
    <a-row type="flex" justify="start" style="padding-top:10px">
      <a-col :span="12" style="text-align:left">
        <a-button
          type="primary"
          style="margin-left:10px"
          @click="
            () => {
              this.addServerVisible = true;
            }
          "
          >Connenct to Redis Server</a-button
        >
      </a-col>
      <a-col :span="12"> </a-col>
    </a-row>
    <a-row
      :style="{ height: windowHeight - 52 + 'px', 'padding-top': 10 + 'px' }"
    >
      <a-col :span="8" class="col-wrap" style="position:relative">
        <a-tree
          class="redis-tree"
          :treeData="treeData"
          :loadData="onLoadData"
          @rightClick="treeRightClick"
          @select="onSelectTree"
        >
        </a-tree>
      </a-col>
      <a-col :span="16" class="col-wrap">
        <a-tabs hideAdd v-model="activeKey">
          <a-tab-pane
            v-for="pane in panes"
            :tab="pane.title"
            :key="pane.key"
            type="editable-card"
            :closable="true"
          >
            <a-row>
              <a-col :span="24" class="key-info">
                <span>HASH:</span>
                <a-input :style="{ width: 200 + 'px' }" />
                <span>TTL:-1</span>
                <a-button type="dashed">rename</a-button>
                <a-button type="danger">delete</a-button>
              </a-col>
            </a-row>
            <a-row>
              <a-col :span="24">
                <a-table
                  size="small"
                  :scroll="{ y: 300 }"
                  :columns="columns"
                  :rowKey="record => record.key"
                  :dataSource="data"
                  bordered
                >
                </a-table>
              </a-col>
            </a-row>
          </a-tab-pane>
        </a-tabs>
      </a-col>
      <a-modal
        title="Connect To Redis Server"
        :visible="addServerVisible"
        @ok="handleOk"
        @cancel="handleCancel"
      >
        <a-form :form="serverForm">
          <a-form-item
            label="Name"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
          >
            <a-input
              v-decorator="[
                'name',
                {
                  rules: [{ required: true, message: 'Please input name!' }]
                }
              ]"
            />
          </a-form-item>
          <a-form-item
            label="Host"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
          >
            <a-input
              v-decorator="[
                'host',
                {
                  initialValue: '127.0.0.1',
                  rules: [{ required: true, message: 'Please input host!' }]
                }
              ]"
            />
          </a-form-item>
          <a-form-item
            label="Port"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
          >
            <a-input
              v-decorator="[
                'port',
                {
                  initialValue: 6379,
                  rules: [{ required: true, message: 'Please input port!' }]
                }
              ]"
            />
          </a-form-item>
          <a-form-item
            label="Password"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
          >
            <a-input type="password" v-decorator="['password']" />
          </a-form-item>
        </a-form>
      </a-modal>
    </a-row>
  </div>
</template>

<script>
import chunk from "lodash/chunk";
import findIndex from "lodash/findIndex";

export default {
  name: "home",
  data: () => {
    return {
      addServerVisible: false,
      windowHeight: document.documentElement.clientHeight,
      treeData: [],
      activeKey: "1",
      columns: [
        { title: "row", dataIndex: "row", width: 100 },
        { title: "key", dataIndex: "key", width: 180 },
        { title: "value", dataIndex: "value", width: 200 }
      ],
      data: [],
      panes: []
    };
  },
  computed: {
    cardList() {
      return chunk(this.dataList, 4);
    }
  },
  beforeCreate() {
    this.serverForm = this.$form.createForm(this);
  },
  methods: {
    handleOk() {
      this.serverForm.validateFields((err, fieldsValue) => {
        if (err) {
          return;
        }
        this.connRedis(fieldsValue);
      });
    },
    async loadRedisServer() {
      // eslint-disable-next-line no-undef
      const redisServers = await loadRedisServer();
      if (redisServers) {
        for (let i = 0; i < redisServers.length; i++) {
          this.treeData.push({
            title: redisServers[i],
            key: redisServers[i],
            level: 0,
            children: []
          });
        }
      }
    },
    async connRedis(connOption) {
      // eslint-disable-next-line no-undef
      const result = await initRedisClient(connOption);
      if (result) {
        this.addServerVisible = false;
        const redisProfile = {
          title: connOption.name,
          key: connOption.name,
          level: 0,
          children: []
        };
        this.treeData.push(redisProfile);
      }
    },
    handleCancel() {
      this.addServerVisible = false;
    },
    async selectDb(name, num) {
      // eslint-disable-next-line no-undef
      const dbKeys = await selectDb(name, num);
      return dbKeys;
    },
    async readKeys(treeNode) {
      debugger;
      const dbKeys = await this.selectDb(
        treeNode.$parent.dataRef.title,
        treeNode.dataRef.value
      );
      for (let i = 0; i < dbKeys.length; i++) {
        treeNode.dataRef.children.push({
          title: dbKeys[i],
          key: dbKeys[i],
          level: 2,
          isLeaf: true
        });
      }
      this.treeData = [...this.treeData];
    },
    onLoadData(treeNode) {
      return new Promise(resolve => {
        switch (treeNode.dataRef.level) {
          case 0:
            for (let i = 0; i < 15; i++) {
              treeNode.dataRef.children.push({
                title: `db${i}`,
                key: `db${i}`,
                value: i,
                level: 1,
                children: []
              });
            }
            this.treeData = [...this.treeData];
            break;
          case 1:
            this.readKeys(treeNode);
            break;
          default:
            break;
        }
        resolve();
      });
    },
    async getValueByKey(treeNode) {
      const key = treeNode.dataRef.key;
      // eslint-disable-next-line no-undef
      const value = await valueByKey(
        treeNode.$parent.$parent.dataRef.title,
        treeNode.$parent.dataRef.value,
        key
      );
      const dataArray = chunk(value, 2);
      this.data = dataArray.map((d, index) => {
        return {
          row: index,
          key: d[0],
          value: d[1]
        };
      });
      const index = findIndex(this.panes, d => d.key == key);
      if (index == -1) {
        this.panes.push({
          title: key,
          key: key
        });
      }
      this.activeKey = key;
      // console.log(value);
    },
    onSelectTree(selectedKeys, e) {
      const treeNode = e.node;
      if (treeNode.dataRef.level === 2) {
        this.getValueByKey(treeNode);
      }
    },
    treeRightClick(e, treeNode) {
      console.log(e);
    }
  },
  mounted() {
    this.loadRedisServer();
  }
};
</script>
<style scoped>
.col-wrap {
  border: 1px solid #000;
  display: flex;
  height: 100%;
}
.redis-tree {
  margin-left: 20px;
}
.key-info span,
.key-info button,
.key-info input {
  margin-left: 10px;
}
</style>
