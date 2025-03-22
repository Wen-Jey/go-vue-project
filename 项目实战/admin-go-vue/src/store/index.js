/**
 * vuex状态管理 
 */

import Vue from 'vue'
import Vuex from 'vuex'
import mutations from './mutations'
import storage from '@/utils/storage'

Vue.use(Vuex)

// 保存各个状态信息集合
const state = new Vuex.Store({
    sysAdmin: "" || storage.getItem("sysAdmin"),
    token: "" || storage.getItem("token"),
    leftMenuList: "" || storage.getItem("leftMenuList"),
    permissionList: "" || storage.getItem("permissionList"),
    activePath: "" || storage.getItem("activePath"),
    mutations
})

export default state