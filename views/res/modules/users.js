layui.define(['admin'], function (exports) {
    'use strict';
    var admin = layui.admin
        , table = layui.table
        , form = layui.form
        , laydate = layui.laydate

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
                        var formattedDate = moment(inputDate).format('Y-MM-DD HH:mm:ss');
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

    exports('users', {});
});