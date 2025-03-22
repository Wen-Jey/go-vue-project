<template>
    <el-card>
        <!--搜索区域-->
        <el-form :inline="true" :model="queryParams">
            <el-form-item label="用户账号" prop="username">
                <el-input size="mini" placeholder="请输入用户账号" clearable v-model="queryParams.username"
                    @keyup.enter.native="handleQuery" />
            </el-form-item>
            <el-form-item label="账号状态" prop="status">
                <el-select size="mini" placeholder="请选择账号状态" v-model="queryParams.status">
                    <el-option v-for="item in statusList" :key="item.value" :label="item.label"
                        :value="item.value"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="开始时间" prop="beginTime">
                <el-date-picker class="input-width" size="mini" type="date" style="width: 190px" value-format="yyyy-MM-dd"
                    clearable placeholder="请选择开始时间" v-model="queryParams.beginTime" @keyup.enter.native="handleQuery" />
            </el-form-item>
            <el-form-item label="结束时间">
                <el-date-picker class="input-width" size="mini" type="date" style="width: 190px" value-format="yyyy-MM-dd"
                    clearable placeholder="请选择结束时间" v-model="queryParams.endTime" @keyup.enter.native="handleQuery" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                <el-button type="primary" icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
            </el-form-item>
            <!--操作按钮-->
            <el-row :gutter="10" class="mb8">
                <el-col :span="1.5">
                    <el-button plain type="primary" icon="el-icon-plus" size="mini"
                        @click="addDialogVisible = true" v-authority="['base:admin:add']">新增</el-button>
                </el-col>
            </el-row>
            <!--列表-->
            <el-table v-loading="Loading" :data="adminList" border stripe style="width: 100%"
                :header-cell-style="{ background: '#eef1f6', color: '#606266' }">
                <el-table-column label="ID" prop="id" v-if="false" />
                <el-table-column label="用户账号" prop="username" />
                <el-table-column label="用户昵称" prop="nickname" />
                <el-table-column label="用户手机" prop="phone" width="120" />
                <el-table-column label="用户邮箱" prop="email" width="170" />
                <el-table-column label="用户头像" prop="icon">
                    <template slot-scope="scope">
                        <el-avatar shape="circle" :src="scope.row.icon"></el-avatar>
                    </template>
                </el-table-column>
                <el-table-column label="创建时间" prop="createTime" width="160" />
                <el-table-column label="账号状态" width="150">
                    <template slot-scope="scope">
                        <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-color="#13ce66"
                            inactive-color="#F5222D" active-text="启用" inactive-text="停用"
                            @change="adminUpdateStatus(scope.row)">
                        </el-switch>
                    </template>
                </el-table-column>
                <el-table-column label="部门/岗位" width="180">
                    <template slot-scope="scope">
                        <div>{{ scope.row.deptName }} / {{ scope.row.postName }}</div>
                    </template>
                </el-table-column>
                <el-table-column label="角色名称" prop="roleName">
                    <template slot-scope="scope">
                        <el-tag type="success">{{ scope.row.roleName }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="个人简介" prop="note" />
                <el-table-column label="更多操作" fixed="right" width="200">
                    <template slot-scope="scope">
                        <el-button size="small" type="text" icon="el-icon-edit"
                            @click="showEditAdminDialog(scope.row.id)" v-authority="['base:admin:edit']">编辑
                        </el-button>
                        <el-button size="small" type="text" icon="el-icon-delete" @click="handleAdminDelete(scope.row)" v-authority="['base:admin:delete']">删除
                        </el-button>
                        <el-button size="small" type="text" icon="el-icon-key" @click="handleResetPwd(scope.row)" v-authority="['base:admin:reset']">重置密码
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!--分页区域-->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryParams.pageNum" :page-sizes="[10, 50, 100, 500, 1000]" :page-size="queryParams.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-form>
        <!-- 添加用户 -->
        <el-dialog title="新增用户" :visible.sync="addDialogVisible" width="40%" @close="addDialogClosed">
            <el-form :model="addForm" :rules="addFormRules" ref="addFormRefForm" label-width="80px">
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户昵称" prop="nickname">
                            <el-input v-model="addForm.nickname" placeholder="请输入用户昵称" maxlength="30" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item size="mini" label="归属部门" prop="deptId">
                            <treeselect v-model="addForm.deptId" :options="deptList" :show-count="true"
                                placeholder="请选择归属部门" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="手机号码" prop="phone">
                            <el-input v-model="addForm.phone" placeholder="请输入手机号码" maxlength="11" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="用户邮箱" prop="email">
                            <el-input v-model="addForm.email" placeholder="请输入邮箱" maxlength="50" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户名称" prop="username">
                            <el-input v-model="addForm.username" placeholder="请输入用户名称" maxlength="30" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="用户密码" prop="password">
                            <el-input v-model="addForm.password" placeholder="请输入用户密码" type="password" maxlength="20"
                                show-password />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户状态" prop="status">
                            <el-radio-group v-model="addForm.status">
                                <el-radio :label="1">正常</el-radio>
                                <el-radio :label="2">停用</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户岗位" prop="postId">
                            <el-select placeholder="请选择岗位" v-model="addForm.postId">
                                <el-option v-for="item in postList" :key="item.id" :label="item.postName" :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="用户角色" prop="roleId">
                            <el-select placeholder="请选择角色" v-model="addForm.roleId">
                                <el-option v-for="item in roleList" :key="item.id" :label="item.roleName"
                                    :value="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="24">
                        <el-form-item label="个人简介" prop="note">
                            <el-input v-model="addForm.note" type="textarea" placeholder="请输入内容"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addAdmin">确 定</el-button>
                <el-button type="primary" @click="addDialogVisible = false">取 消</el-button>
            </span>
        </el-dialog>
        <!--修改用户-->
        <el-dialog title="修改用户" :visible.sync="editDialogVisible" width="40%" @close="editDialogClosed">
            <el-form :model="adminInfo" :rules="editFormRules" ref="editFormRefForm" label-width="80px">
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户昵称" prop="nickname">
                            <el-input v-model="adminInfo.nickname" placeholder="请输入用户昵称" maxlength="30" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item size="mini" label="归属部门" prop="deptId">
                            <treeselect v-model="adminInfo.deptId" :options="deptList" :show-count="true"
                                placeholder="请选择归属部门" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="手机号码" prop="phone">
                            <el-input v-model="adminInfo.phone" placeholder="请输入手机号码" maxlength="11" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="用户邮箱" prop="email">
                            <el-input v-model="adminInfo.email" placeholder="请输入邮箱" maxlength="50" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户名称" prop="username">
                            <el-input v-model="adminInfo.username" placeholder="请输入用户名称" maxlength="30" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="用户状态" prop="status">
                            <el-radio-group v-model="adminInfo.status">
                                <el-radio :label="1">正常</el-radio>
                                <el-radio :label="2">停用</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="用户岗位" prop="postId">
                            <el-select placeholder="请选择岗位" v-model="adminInfo.postId">
                                <el-option v-for="item in postList" :key="item.id" :label="item.postName" :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="用户角色" prop="roleId">
                            <el-select placeholder="请选择角色" v-model="adminInfo.roleId">
                                <el-option v-for="item in roleList" :key="item.id" :label="item.roleName"
                                    :value="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="24">
                        <el-form-item label="个人简介" prop="note">
                            <el-input v-model="adminInfo.note" type="textarea" placeholder="请输入内容"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editAdminInfo">确 定</el-button>
                <el-button type="primary" @click="editDialogVisible = false">取 消</el-button>
            </span>
        </el-dialog>
    </el-card>
</template>

<script>
import Treeselect from "@riophae/vue-treeselect"
import "@riophae/vue-treeselect/dist/vue-treeselect.css"

export default {
    components: { Treeselect },
    data() {
        return {
            statusList: [{
                value: '1',
                label: '启用',
            }, {
                value: '2',
                label: '停用',
            }],
            Loading: false,
            queryParams: {},
            adminList: [],
            total: 0,
            addDialogVisible: false,
            deptList: [],
            roleList: [],
            postList: [],
            addForm: {
                username: '',
                password: '',
                deptId: undefined,
                postId: undefined,
                roleId: undefined,
                email: '',
                nickname: '',
                status: 1,
                phone: '',
                note: ''
            },
            addFormRules: {
                deptId: [{ required: true, message: '请选择部门', trigger: 'blur' }],
                postId: [{ required: true, message: '请选择岗位', trigger: 'blur' }],
                roleId: [{ required: true, message: '请选择角色', trigger: 'blur' }],
                username: [{ required: true, message: '请输入用户账号', trigger: 'blur' }],
                password: [{ required: true, message: '请输入用户密码', trigger: 'blur' }],
                status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
                email: [{ required: true, message: '请输入用户邮箱', trigger: 'blur' }],
                nickname: [{ required: true, message: '请输入用户昵称', trigger: 'blur' }],
                phone: [{ required: true, message: '请输入用户手机', trigger: 'blur' }]
            },
            editDialogVisible: false,
            adminInfo: {},
            editFormRules: {
                deptId: [{ required: true, message: '请选择部门', trigger: 'blur' }],
                postId: [{ required: true, message: '请选择岗位', trigger: 'blur' }],
                roleId: [{ required: true, message: '请选择角色', trigger: 'blur' }],
                username: [{ required: true, message: '请输入用户账号', trigger: 'blur' }],
                status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
                email: [{ required: true, message: '请输入用户邮箱', trigger: 'blur' }],
                nickname: [{ required: true, message: '请输入用户昵称', trigger: 'blur' }],
                phone: [{ required: true, message: '请输入用户手机', trigger: 'blur' }]
            },

        }
    },
    methods: {
        // 查询列表
        async getAdminList() {
            this.Loading = true
            const { data: res } = await this.$api.queryAdminList(this.queryParams)
            if (res.code !== 200) {
                this.$message.error(res.message)
            } else {
                this.adminList = res.data.list
                this.total = res.data.total
                this.Loading = false
            }
        },
        // 搜索按钮操作
        handleQuery() {
            this.getAdminList();
        },
        // 重置按钮操作
        resetQuery() {
            this.queryParams = {}
            this.getAdminList();
            this.$message.success("重置成功")
        },
        // pageSize
        handleSizeChange(newSize) {
            this.queryParams.pageSize = newSize
            this.getAdminList()
        },
        // pageNum
        handleCurrentChange(newPage) {
            this.queryParams.pageNum = newPage
            this.getAdminList()
        },
        // 修改用户状态
        async adminUpdateStatus(row) {
            let text = row.status === 2 ? "停用" : "启用";
            const confirmResult = await this.$confirm('确认要"' + text + '""' + row.username + '"用户吗?', "警告", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning",
            }).catch(err => err)
            if (confirmResult !== 'confirm') {
                await this.getAdminList()
                return this.$message.info('已取消删除')
            }
            this.$api.updateAdminStatus(row.id, row.status)
            this.$message.success(text + "成功")
            await this.getAdminList()
        },
        // 部门下拉列表
        async getDeptVoList() {
            const { data: res } = await this.$api.querySysDeptVoList()
            if (res.code !== 200) {
                this.$message.error(res.message)
            } else {
                this.deptList = this.$handleTree.handleTree(res.data, "id");
            }
        },
        // 角色下拉列表
        async getRoleVoList() {
            const { data: res } = await this.$api.querySysRoleVoList()
            if (res.code !== 200) {
                this.$message.error(res.message)
            } else {
                this.roleList = res.data
            }
        },
        // 岗位下拉列表
        async getPostVoList() {
            const { data: res } = await this.$api.querySysPostVoList()
            if (res.code !== 200) {
                this.$message.error(res.message)
            } else {
                this.postList = res.data
            }
        },
        // 监听添加用户对话框关闭
        addDialogClosed() {
            this.$refs.addFormRefForm.resetFields()
        },
        // 新增用户
        addAdmin() {
            this.$refs.addFormRefForm.validate(async valid => {
                if (!valid) return
                const { data: res } = await this.$api.addAdmin(this.addForm);
                if (res.code !== 200) {
                    this.$message.error(res.message)
                } else {
                    this.$message.success('新增用户成功')
                    this.addDialogVisible = false
                    await this.getAdminList()
                }
            })
        },
        // 展示编辑用户的对话框
        async showEditAdminDialog(id) {
            const { data: res } = await this.$api.adminInfo(id)
            if (res.code !== 200) {
                this.$message.error(res.message)
            } else {
                this.adminInfo = res.data
                this.editDialogVisible = true
            }
        },
        // 监听修改用户对话框的关闭事件
        editDialogClosed() {
            this.$refs.editFormRefForm.resetFields()
        },
        // 修改用户信息并提交
        editAdminInfo() {
            this.$refs.editFormRefForm.validate(async valid => {
                if (!valid) return
                const { data: res } = await this.$api.adminUpdate(this.adminInfo)
                if (res.code !== 200) {
                    this.$message.error(res.message)
                } else {
                    this.editDialogVisible = false
                    await this.getAdminList()
                    this.$message.success('修改用户成功')
                }
            })
        },
        // 删除
        async handleAdminDelete(row) {
            const confirmResult = await this.$confirm('是否确认删除用户为"' + row.username + '"的数据项？', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).catch(err => err)
            if (confirmResult !== 'confirm') {
                return this.$message.info('已取消删除')
            }
            const { data: res } = await this.$api.deleteAdmin(row.id)
            if (res.code !== 200) {
                this.$message.error(res.message)
            } else {
                this.$message.success('删除成功')
                await this.getAdminList()
            }
        },
        // 重置密码
        handleResetPwd(row) {
            this.$prompt('请输入"' + row.username + '"的新密码', "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                closeOnClickModal: false,
                inputPattern: /^.{5,20}$/,
                inputErrorMessage: "用户密码长度必须介于 5 和 20 之间"
            }).then(({ value }) => {
                this.$api.resetPassword(row.id, value).then(response => {
                    this.$message.success("修改成功，新密码是：" + value);
                });
            }).catch(() => {
            });
        }
    },
    created() {
        this.getAdminList()
        this.getDeptVoList()
        this.getRoleVoList()
        this.getPostVoList()
    }
}
</script>

<style lang="less" scoped></style>