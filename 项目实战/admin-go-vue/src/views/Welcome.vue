<template>
    <el-row>
        <el-col>
            <el-card>
                <div ref="lineChart" style="height: 380px; width: 1200px; margin: 0px auto;"></div>
            </el-card>
            <div class="graph">
                <el-card style="height: 360px;">
                    <div ref="Histogram" style="height: 380px; width: 600px; margin: 0px auto;"></div>
                </el-card>
                <el-card style="height: 360px;">
                    <div ref="PieChart" style="height: 380px; width: 600px; margin: 0px auto;"></div>
                </el-card>
            </div>
        </el-col>
    </el-row>
</template>

<script>
import * as echarts from 'echarts'
export default {
    data() {
        return {}
    },
    mounted() {
        // 折线图
        const lineChart = echarts.init(this.$refs.lineChart)
        var optionLineChart = {
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'shadow'
                }
            },
            legend: {},
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: 'Direct',
                    type: 'bar',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [320, 332, 301, 334, 390, 330, 320]
                },
                {
                    name: 'Email',
                    type: 'bar',
                    stack: 'Ad',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [120, 132, 101, 134, 90, 230, 210]
                },
                {
                    name: 'Union Ads',
                    type: 'bar',
                    stack: 'Ad',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [220, 182, 191, 234, 290, 330, 310]
                },
                {
                    name: 'Video Ads',
                    type: 'bar',
                    stack: 'Ad',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [150, 232, 201, 154, 190, 330, 410]
                },
                {
                    name: 'Search Engine',
                    type: 'bar',
                    data: [862, 1018, 964, 1026, 1679, 1600, 1570],
                    emphasis: {
                        focus: 'series'
                    },
                    markLine: {
                        lineStyle: {
                            type: 'dashed'
                        },
                        data: [[{ type: 'min' }, { type: 'max' }]]
                    }
                },
                {
                    name: 'Baidu',
                    type: 'bar',
                    barWidth: 5,
                    stack: 'Search Engine',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [620, 732, 701, 734, 1090, 1130, 1120]
                },
                {
                    name: 'Google',
                    type: 'bar',
                    stack: 'Search Engine',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [120, 132, 101, 134, 290, 230, 220]
                },
                {
                    name: 'Bing',
                    type: 'bar',
                    stack: 'Search Engine',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [60, 72, 71, 74, 190, 130, 110]
                },
                {
                    name: 'Others',
                    type: 'bar',
                    stack: 'Search Engine',
                    emphasis: {
                        focus: 'series'
                    },
                    data: [62, 82, 91, 84, 109, 110, 120]
                }
            ]
        };
        lineChart.setOption(optionLineChart)

        // 柱状图
        const Histogram = echarts.init(this.$refs.Histogram)
        var optionHistogram = {
            xAxis: {
                type: 'category',
                data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
            },
            yAxis: {
                type: 'value'
            },
            series: [
                {
                    data: [120, 200, 150, 80, 70, 110, 130],
                    type: 'bar'
                }
            ]
        };
        Histogram.setOption(optionHistogram)

        // 饼状图
        const PieChart = echarts.init(this.$refs.PieChart)
        var optionPieChart = {
            tooltip: {
                trigger: 'item'
            },
            legend: {
                top: '5%',
                left: 'center'
            },
            series: [
                {
                    name: 'Access From',
                    type: 'pie',
                    radius: ['40%', '70%'],
                    avoidLabelOverlap: false,
                    label: {
                        show: false,
                        position: 'center'
                    },
                    emphasis: {
                        label: {
                            show: true,
                            fontSize: 40,
                            fontWeight: 'bold'
                        }
                    },
                    labelLine: {
                        show: false
                    },
                    data: [
                        { value: 1048, name: 'Search Engine' },
                        { value: 735, name: 'Direct' },
                        { value: 580, name: 'Email' },
                        { value: 484, name: 'Union Ads' },
                        { value: 300, name: 'Video Ads' }
                    ]
                }
            ]
        };
        PieChart.setOption(optionPieChart)

    }
};
</script>

<style lang="less" scoped>
.graph {
    margin-top: 20px;
    display: flex;
    justify-content: space-between;

    .el-card {
        width: 48%;
    }
}
</style>