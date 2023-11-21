layui.define(['admin', 'table', 'form'], function (exports) {
    'use strict';
    var admin = layui.admin
        , table = layui.table
        , form = layui.form;

    form.render();

    table.render({
        elem: '#KYO-provinces-table'
        , url: '/api/table/provinces/get'
        , cols: [[
            { field: 'id', width: 100, title: 'ID' }
            , { field: 'province_id', title: '省ID' }
            , { field: 'province_name', title: '省名称' }
        ]]
    });

    // 搜索提交
    form.on('submit(KYO-provinces-table-search)', function (obj) {
        var field = obj.field; // 获得表单字段
        // 执行搜索重载
        table.reload('KYO-provinces-table', {
            // page: {
            //     curr: 1 // 重新从第 1 页开始
            // },
            where: field // 搜索的字段
        });
        return false; // 阻止默认 form 跳转
    });

    exports('provinces', {})
});