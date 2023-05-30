import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { Button, Form, Field, Pagination,List, Cell, CellGroup, Col, Row  } from 'vant';

import 'vant/lib/index.css';

createApp(App)
.use(Button)
.use(Form)
.use(Field)
.use(CellGroup)
.use(Pagination)
.use(List)
.use(Cell)
.use(Col)
.use(Row)
.use(router)
.mount('#app');
