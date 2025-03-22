<template>
    <div class="tags">
        <el-tag class="tag" size="medium" closable :effect="item.title == $route.meta.tTitle ? 'dark' : 'plain'"
            v-for="item, index in tags" :key="item.path" @click="goTo(item.path)" @close="close(index)"
            :closable="index > 0" :disable-transitions="true">
            <i class="circular" v-show="item.title == $route.meta.tTitle"></i>
            {{ item.title }}</el-tag>
    </div>
</template>

<script>
export default {
    name: "Tags",
    data() {
        return {
            tags: [{
                title: "首页",
                path: "/welcome"
            }]
        }
    },
    watch: {
        $route: {
            immediate: true,
            handler(val) {
                const boolean = this.tags.find(item => {
                    return val.path == item.path
                })
                if (!boolean) {
                    this.tags.push({
                        title: val.meta.tTitle,
                        path: val.path
                    })
                }
            }
        }
    },
    methods: {
        // 点击跳转到指定位置
        goTo(path) {
            this.$router.push(path)
        },
        // 点击关闭
        close(i) {
            if (this.tags[i].path == this.$route.meta.path && i !== this.tags.length - 1) {
                this.$router.push(this.tags[this.tags.length - 1].path)
            } else if (i === this.tags.length - 1) {
                this.$router.push(this.tags[this.tags.length - 2].path)
            }
            this.tags.splice(i, 1)
        }
    }
}
</script>

<style lang="less" scoped>
.tags {
    padding-left: 20px;
    padding-top: 2px;
    padding-bottom: 2px;

    .tag {
        cursor: pointer;
        margin-right: 3px;

        .circular {
            width: 8px;
            height: 8px;
            margin-right: 4px;
            background-color: #fff;
            border-radius: 50%;
            display: inline-block;
        }
    }
}
</style>