import React,{PropTypes} from 'react'
import {Table,Popconfirm} from 'antd'


const ArticleList = ({list})=>{

   

    const columns = [{
        title:'标题',
        dataIndex:'title',
        key:'title'
    },{
       title:'内容',
         dataIndex:'context',
        key:'context'
    },
    {
        title: '操作',
        key: 'operation',
        render: (text, record) => (
        <p>
            <a onClick={() => {}}>编辑</a>
            &nbsp;
            <Popconfirm title="确定要删除吗？" onConfirm={() => {}}>
            <a>删除</a>
            </Popconfirm>
        </p>
        ),
  }];

  return (
      <div>
        <Table 
            columns={columns}
             dataSource={list}
        />
      </div>
  )
}

// PriceList.propTypes = {
//      list: PropTypes.array,
// }

export default ArticleList