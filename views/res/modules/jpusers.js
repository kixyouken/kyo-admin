layui.define(['admin'], function (exports) {
    'use strict';
    var admin = layui.admin
        , table = layui.table
        , form = layui.form
        , laydate = layui.laydate
        , $ = layui.$
        , setter = layui.setter

    var localStorage = layui.data(setter.tableName);

    laydate.render({
        elem: '#KYO-jpusers-birthday'
    })

    form.render();

    table.render({
        elem: '#KYO-jpusers-table'
        , url: '/api/table/users/paginate'
        , cols: [[
            { field: 'id', width: 100, title: 'ID' }
            , { field: 'name', title: '姓名' }
            , {
                field: 'birthday', title: '生日', templet: function (d) {
                    if (d.birthday != null) {
                        // 将日期字符串转换为Date对象
                        var inputDate = new Date(d.birthday);
                        // 使用Moment.js进行日期格式化
                        var formattedDate = moment(inputDate).format('Y-MM-DD');
                        return formattedDate;
                    } else {
                        return "";
                    }
                }
            }
            , { field: 'mobile', title: '电话' }
            , { field: 'email', title: '邮箱' }
            , { field: 'master_university_name', title: '学校名' }
            , {
                field: 'created_at', title: '创建时间', templet: function (d) {
                    if (d.created_at != null) {
                        // 将日期字符串转换为Date对象
                        var inputDate = new Date(d.created_at);
                        // 使用Moment.js进行日期格式化
                        var formattedDate = moment(inputDate).format('Y-MM-DD HH:mm:ss');
                        return formattedDate;
                    } else {
                        return "";
                    }
                }
            }
            , { title: '操作', width: 150, toolbar: '#KYO-jpusers-table-bar' }
        ]]
        , page: true
    });

    table.on('tool(KYO-jpuser-table)', function (obj) {
        var data = obj.data;
        console.log(data);
        switch (obj.event) {
            case 'view':
                admin.popup({
                    id: 'KYO-user-table-view-popup' //定义唯一ID，防止重复弹出
                    , area: ['80%', '80%']
                    , success: function () {
                        //将 views 目录下的某视图文件内容渲染给该面板
                        layui.view(this.id).render('jpusers/view', data);
                    }
                });
                break;
            case 'edit':
                admin.popup({
                    id: 'KYO-user-table-edit-popup' //定义唯一ID，防止重复弹出
                    , area: ['80%', '80%']
                    , success: function () {
                        // 将 views 目录下的某视图文件内容渲染给该面板
                        layui.view(this.id).render('jpusers/edit', data);
                    }
                });
                break;
            case 'delete':
                layer.confirm('确定删除吗？', { icon: 2 }, function () {
                    admin.req({
                        url: '/api/table/users/' + data.user_id + '?access_token=' + localStorage[setter.request.tokenName]
                        , type: 'delete'
                        , done: function (res) {
                            layer.msg(res.msg);
                            layer.closeLast('dialog');
                            table.reloadData('KYO-users-table');
                        }
                    })
                })
                break;
            default:
                break;
        }
    })

    // 搜索提交
    form.on('submit(KYO-jpusers-table-search)', function (obj) {
        var field = obj.field; // 获得表单字段
        // 执行搜索重载
        table.reload('KYO-jpusers-table', {
            page: {
                curr: 1 // 重新从第 1 页开始
            },
            where: field // 搜索的字段
        });
        return false; // 阻止默认 form 跳转
    });

    exports('jpusers', {});
});