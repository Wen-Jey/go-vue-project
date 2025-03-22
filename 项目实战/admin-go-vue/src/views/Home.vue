<template>
    <el-container class="home-container">
        <el-aside :width="isCollapse ? '64px' : '200px'">
            <div class="logo">
                <img src="../assets/image/logo.jpg" class="siderbar-logo">
                <h3 v-show="!isCollapse">管理系统(项目实战)</h3>
            </div>
            <el-menu background-color="#304156" text-color="#fff" unique-opened router :default-active="$route.path"
                :collapse="isCollapse" :collapse-transition="false">
                <!--无子集菜单-->
                <el-menu-item :index="'/' + item.url" v-for="item in noChildren" :key="item.menuName"
                    @click="saveNavState('/' + item.url)">
                    <i :class="item.icon"></i>
                    <template slot="title">
                        <span>{{ item.menuName }}</span>
                    </template>
                </el-menu-item>
                <!--有子集菜单-->
                <el-submenu :index="item.id + ''" v-for="item in hasChildren" :key="item.id">
                    <template slot="title">
                        <i :class="item.icon"></i>
                        <span>{{ item.menuName }}</span>
                    </template>
                    <el-menu-item :index="'/' + subItem.url" v-for="subItem in item.menuSvoList" :key="subItem.id"
                        @click="saveNavState('/' + subItem.url)">
                        <template slot="title">
                            <i :class="subItem.icon"></i>
                            <span>{{ subItem.menuName }}</span>
                        </template>
                    </el-menu-item>
                </el-submenu>
            </el-menu>
        </el-aside>
        <el-container>
            <el-header>
                <div class="fold-btn">
                    <i :class="collapseBtnClass" @click="toggleCollapse"></i>
                </div>
                <div class="bread-btn">
                    <el-breadcrumb separator="/" v-if="$router.currentRoute.path != '/welcome'">
                        <el-breadcrumb-item :to="{ path: '/welcome' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>{{ $route.meta.sTitle }}</el-breadcrumb-item>
                        <el-breadcrumb-item>{{ $route.meta.tTitle }}</el-breadcrumb-item>
                    </el-breadcrumb>
                    <el-breadcrumb separator="/" v-else>
                        <el-breadcrumb-item :to="{ path: '/welcome' }">首页</el-breadcrumb-item>
                    </el-breadcrumb>
                </div>
                <HeadImage />
            </el-header>
            <Tags />
            <el-main><router-view /></el-main>
        </el-container>
    </el-container>
</template>

<script>
import storage from '@/utils/storage'
import HeadImage from '@/components/HeadImage'
import Tags from '@/components/Tags'

export default {
    name: "Home",
    components: { HeadImage, Tags },
    data() {
        return {
            leftMenuList: storage.getItem("leftMenuList"),
            activePath: '',
            collapseBtnClass: "el-icon-s-fold",
            isCollapse: false,

        }
    },
    computed: {
        // 无子集
        noChildren() {
            return this.leftMenuList.filter(item => !item.menuSvoList)
        },
        // 有子集
        hasChildren() {
            return this.leftMenuList.filter(item => item.menuSvoList)
        }
    },
    methods: {
        // 点击实现跳转
        saveNavState(activePath) {
            storage.setItem('activePath', activePath)
            this.activePath = activePath
        },
        // 张开和折叠
        toggleCollapse() {
            this.isCollapse = !this.isCollapse
            if (this.isCollapse) {
                this.collapseBtnClass = 'el-icon-s-unfold'
            } else {
                this.collapseBtnClass = 'el-icon-s-fold'
            }
        }
    }
}
</script>

<style lang="less" scoped>
.home-container {
    height: 100%;

    .el-aside {
        background-color: #304156;

        .logo {
            margin-top: 5px;
            display: flex;
            align-items: center;
            font-size: 13px;
            height: 50px;
            color: #fff;
            font-style: italic;

            .siderbar-logo {
                width: 32px;
                height: 32px;
                margin: 0 16px;
            }
        }

        .el-menu {
            border-right: none;
        }
    }

    .el-header {
        background-color: #f9fafc;
        align-items: center;
        justify-content: space-between;
        display: flex;

        .fold-btn {
            padding-top: 2px;
            font-size: 23px;
            cursor: pointer;
        }

        .bread-btn {
            padding-top: 2px;
            position: fixed;
            margin-left: 40px;
        }
    }

    .el-main {
        background-color: #eaedf1;
    }
}
</style>