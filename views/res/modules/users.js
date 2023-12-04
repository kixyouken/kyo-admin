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
        elem: '#KYO-users-birthday'
        , range: true
    })

    form.render();

    table.render({
        elem: '#KYO-users-table'
        , url: '/api/table/users/paginate'
        , cols: [[
            { field: 'user_id', width: 100, title: 'ID' }
            , { field: 'user_name', title: '姓名' }
            , {
                field: '', title: '生日', templet: function (d) {
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
            , { field: 'provinces_province_name', title: '省' }
            , { field: 'citys_city_name', title: '市' }
            , { field: 'countys_county_name', title: '县' }
            , {
                field: 'state', title: '状态', templet: function (d) {
                    if (d.state == 1) {
                        return '删除';
                    } else {
                        return '正常';
                    }
                }
            }
            , { title: '操作', width: 150, toolbar: '#KYO-users-table-bar' }
        ]]
        , page: true
    });

    // 搜索提交
    form.on('submit(KYO-users-table-search)', function (obj) {
        var field = obj.field; // 获得表单字段
        // 执行搜索重载
        table.reload('KYO-users-table', {
            page: {
                curr: 1 // 重新从第 1 页开始
            },
            where: field // 搜索的字段
        });
        return false; // 阻止默认 form 跳转
    });

    table.on('tool(KYO-user-table)', function (obj) {
        var data = obj.data;
        switch (obj.event) {
            case 'view':
                admin.popup({
                    id: 'KYO-user-table-view-popup' //定义唯一ID，防止重复弹出
                    , area: ['80%', '80%']
                    , success: function () {
                        //将 views 目录下的某视图文件内容渲染给该面板
                        layui.view(this.id).render('users/view', data);
                    }
                });
                break;
            case 'edit':
                admin.popup({
                    id: 'KYO-user-table-edit-popup' //定义唯一ID，防止重复弹出
                    , area: ['80%', '80%']
                    , success: function () {
                        // 将 views 目录下的某视图文件内容渲染给该面板
                        layui.view(this.id).render('users/edit', data);
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

    form.on('submit(KYO-user-save)', function (obj) {
        var id = obj.field.user_id;
        delete obj.field.user_id;
        admin.req({
            headers: {
                'Content-Type': 'application/json',
            },
            url: '/api/form/users/' + id + '?access_token=' + localStorage[setter.request.tokenName]
            , type: 'put'
            , data: JSON.stringify(obj.field)
            , done: function (res) {
                layer.msg(res.msg);
                layer.closeLast('page');
                table.reloadData('KYO-users-table');
            }
        })
        return false
    })

    exports('users', {});
});