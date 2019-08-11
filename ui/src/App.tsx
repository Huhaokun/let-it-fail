import React, {Component} from 'react';
import Button from 'antd/es/button';
import './App.css';
import {Table} from "antd";

const dataSource = [
    {
        id: 'abc',
        host: '127.0.0.1',
    },
    {
        id: 'edf',
        host: '127.0.0.2',
    }
];

const columns = [
    {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
    },
    {
        title: 'IP',
        dataIndex: 'host',
        key: 'host',
    }
];

class App extends Component {
    render() {
        return (
            <div className="App">
                <Table dataSource={dataSource} columns={columns} />
            </div>
        )
    }
}

export default App;
