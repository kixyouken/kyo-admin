layui.define(['admin'], function (exports) {
    'use strict';
    var admin = layui.admin

    layui.use(['carousel', 'echarts'], function () {
        var $ = layui.$
            , carousel = layui.carousel
            , echarts = layui.echarts;

        admin.req({
            url: '/api/echarts/logins/line'
            , done: function (res) {
                console.log(res);
            }
        })

        var echartsApp = [], options = [
            {
                tooltip: {
                    trigger: 'axis'
                },
                calculable: true,
                legend: {
                    data: ['访问量']
                },

                xAxis: [
                    {
                        type: 'category',
                        data: ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月']
                    }
                ],
                yAxis: [
                    {
                        type: 'value',
                        name: '访问量',
                        axisLabel: {
                            formatter: '{value} 万'
                        }
                    }
                ],
                series: [
                    {
                        name: '访问量',
                        type: 'line',
                        data: [1900, 850, 950, 1000, 1100, 1050, 1000, 1150, 1250, 1370, 1250, 1100]
                    }
                ]
            }
        ]
            , elemDataView = $('#LAY-index-logins').children('div')
            , renderDataView = function (index) {
                echartsApp[index] = echarts.init(elemDataView[index], layui.echartsTheme);
                echartsApp[index].setOption(options[index]);
                window.onresize = echartsApp[index].resize;
            };
        //没找到DOM，终止执行
        if (!elemDataView[0]) return;
        renderDataView(0);

    });

    exports('logins_line', {});
});