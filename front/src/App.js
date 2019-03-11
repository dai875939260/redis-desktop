import React, { Component } from 'react';
import { Button,Row, Col, Tree, Card, Tabs, Input, Table, Modal, Spin } from 'antd';
import chunk from 'lodash/chunk';
import findIndex from 'lodash/findIndex';
import './App.css';

const { TreeNode } = Tree;
const { TabPane } = Tabs;
const { Column } = Table;

class App extends Component {

  constructor(props){
    super();

    this.state = {
      addServerVisible: false,
      windowHeight: document.documentElement.clientHeight,
      treeData: [],
      activeKey: "1",
      tabLoading: false,
      columns: [
        { title: "row", dataIndex: "row", width: 100 },
        { title: "key", dataIndex: "key", width: 180 },
        { title: "value", dataIndex: "value", width: 200 }
      ],
      data: [],
      panes: []
    }
  }

  loadRedisServer = async () => {
    // eslint-disable-next-line no-undef
    const redisServers = await loadRedisServer();
    if (redisServers) {
      const { treeData } = this.state;
      for (let i = 0; i < redisServers.length; i++) {
        treeData.push({
          title: redisServers[i],
          key: redisServers[i],
          level: 0,
          children: []
        });
      }
      this.setState({
        treeData
      })
    }
  }

  onLoadData = treeNode => new Promise(resolve => {
    const { dataRef } = treeNode.props;
    switch (dataRef.level) {
      case 0:
        for (let i = 0; i < 15; i++) {
          dataRef.children.push({
            title: `db${i}`,
            key: `${dataRef.key}-db${i}`,
            value: i,
            level: 1,
            children: []
          });
        }
        this.setState({
          treeData: [...this.state.treeData],
        });
        break;
      case 1:
        this.readKeys(treeNode);
        break;
      default:
        break;
    }
    resolve();
  })

  selectDb = async (name, num) => {
    // eslint-disable-next-line no-undef
    const dbKeys = await selectDb(name, num);
    return dbKeys;
  }

  readKeys = async (treeNode) => {
    const { dataRef } = treeNode.props;
    const { key } = dataRef;
    const keyArray = key.split('-');
    const dbKeys = await this.selectDb(
      keyArray[0],
      dataRef.value
    );
    for (let i = 0; i < dbKeys.length; i++) {
      dataRef.children.push({
        title: dbKeys[i],
        key: key +'-'+ dbKeys[i],
        level: 2,
        isLeaf: true
      });
    }
    this.setState({
      treeData: [...this.state.treeData],
    });
  }

  onSelectTree = (selectedKeys, e) => {
    this.setState({
      tabLoading: true
    })
    const treeNode = e.node;
    if (treeNode.props.dataRef.level === 2) {
      this.getValueByKey(treeNode);
    }
    this.setState({
      tabLoading: false
    })
  }

  onTabChange = activeKey => {
    this.setState({
      activeKey
    })
  }

  onTabEdit = targetKey => {
    const { panes } = this.state;
    const index = findIndex(panes, { key: targetKey});
    console.log(index);
    if(index > -1){
      panes.splice(index, 1);
      let activeKey = "";
      if(panes.length > 0){
        activeKey = panes[0].key;
      }
      this.setState({
        panes,
        activeKey
      })
    }
  }

  handleOk = () => {
    this.setState({
      addServerVisible: false
    })
  }

  handleCancel = () => {
    this.setState({
      addServerVisible: false
    })
  }

  getValueByKey = async treeNode => {
    const { key, title } = treeNode.props.dataRef;
    const { activeKey, panes } = this.state;
    if(activeKey === key){
      return;
    }
    const keyArray = key.split('-');
    // eslint-disable-next-line no-undef
    const storeData = await valueByKey(
      keyArray[0],
      parseInt(keyArray[1].substring(2)),
      keyArray[2]
    );
    const { keyType, value } = storeData;
    const dataArray = chunk(value, 2);
    const data = dataArray.map((d, index) => {
      return {
        row: index,
        key: d[0],
        value: d[1]
      };
    });
    const index = findIndex(panes, {key});
    if (index === -1) {
      panes.push({
        title,
        key,
        keyType
      });
    }
    this.setState({
      panes,
      data,
      activeKey: key
    });
  }

  renderTreeNodes = data => data.map((item) => {
    if (item.children) {
      return (
        <TreeNode title={item.title} key={item.key} dataRef={item}>
          {this.renderTreeNodes(item.children)}
        </TreeNode>
      );
    }
    return <TreeNode {...item} dataRef={item} />;
  })

  componentDidMount(){
    this.loadRedisServer();
  }

  render() {
    const { panes, data, addServerVisible, activeKey, tabLoading } = this.state;
    return (
      <div className="App">
        <Row gutter={16} style={{paddingTop:10}}>
          <Col className="gutter-row" span={12}>
            <Button type="primary" style={{marginLeft:10}} onClick={() => {
              this.setState({
                addServerVisible: true
              })
            }}>Connenct to Redis Server</Button>
          </Col>
        </Row>
        <Row gutter={16} style={{paddingTop:10}}>
          <Col className="gutter-row col-wrap" span={8}>
            <Card style={{marginLeft: 10}}>
              <Tree loadData={this.onLoadData} onSelect={this.onSelectTree}>
                {this.renderTreeNodes(this.state.treeData)}
              </Tree>
            </Card>
          </Col>
          <Col className="gutter-row col-wrap" span={16}>
            <Spin spinning={tabLoading}>
              <Card>
                <Tabs hideAdd 
                  type="editable-card"
                  activeKey={activeKey}
                  onChange={this.onTabChange}
                  onEdit={this.onTabEdit}>
                  {
                    panes.map(p => {
                      return (
                          <TabPane tab={p.title} key={p.key}>
                          <Row>
                            <Col span={24} className="key-info">
                              <span>{p.keyType}:</span>
                              <Input style={{width: 200}}/>
                              <span>TTL:-1</span>
                              <Button type="dashed">rename</Button>
                              <Button type="danger">delete</Button>
                            </Col>
                          </Row>
                          <Row style={{paddingTop: 10}}>
                            <Col span={24}>
                              <Table 
                                size="small"
                                scroll={{ y: 300 }}
                                dataSource={data}
                                pagination={false}
                                bordered>
                                <Column 
                                  title="row"
                                  dataIndex="row"
                                  width="10%"
                                  key="row" />
                                <Column 
                                  title="key"
                                  dataIndex="key"
                                  width="30%"
                                  key="key" />
                                <Column 
                                  title="value"
                                  dataIndex="value"
                                  render={(text) => {
                                    return (
                                      <div style={{ wordWrap: 'break-word', wordBreak: 'break-all' }}>
                                        {text}
                                      </div>
                                    )
                                  }}
                                  width="60%"
                                  key="value" />
                              </Table>
                            </Col>
                          </Row>
                        </TabPane>
                      )
                    })
                  }
                </Tabs>
              </Card>
            </Spin>
          </Col>
        </Row>
        <Modal 
          title="Connect To Redis Server"
          visible={addServerVisible}
          onOk={this.handleOk}
          onCancel={this.handleCancel}
        />
      </div>
    );
  }
}

export default App;
