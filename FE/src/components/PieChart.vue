<template>
    <div class="w-full h-64">
        <canvas ref="chartCanvas"></canvas>
    </div>
</template>

<script>
import { Chart, registerables } from 'chart.js';
Chart.register(...registerables);

export default {
    name: 'PieChart',
    props: {
        data: {
            type: Array,
            required: true
        }
    },
    mounted() {
        this.renderChart();
    },
    watch: {
        data: {
            handler() {
                this.renderChart();
            },
            deep: true
        }
    },
    methods: {
        renderChart() {
            if (this.chart) {
                this.chart.destroy();
            }

            const ctx = this.$refs.chartCanvas.getContext('2d');
            this.chart = new Chart(ctx, {
                type: 'pie',
                data: {
                    labels: this.data.map(item => item.name),
                    datasets: [
                        {
                            label: 'Data',
                            data: this.data.map(item => item.value),
                            backgroundColor: [
                                '#EF4444', // merah
                                '#3B82F6', // biru
                                '#F59E0B', // kuning
                                '#10B981', // hijau
                                '#8B5CF6', // ungu
                                '#F43F5E', // pink
                                '#22D3EE', // cyan
                                '#6366F1'  // indigo
                            ]
                        }
                    ]
                },
                options: {
                    responsive: true,
                    plugins: {
                        legend: {
                            position: 'bottom'
                        }
                    }
                }
            });
        }
    },
    beforeUnmount() {
        if (this.chart) {
            this.chart.destroy();
        }
    }
};
</script>

<style scoped>
canvas {
    max-height: 100%;
}
</style>