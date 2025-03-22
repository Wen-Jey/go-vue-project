/**
 * vue使用配置
 */
import authority from "./Authority"
import Vue from "vue"

const install = function (Vue) {
    Vue.directive('authority', authority)
}

if (window.Vue) {
    window.authority = authority
    Vue.use(install)
}

authority.install = install
export default authority