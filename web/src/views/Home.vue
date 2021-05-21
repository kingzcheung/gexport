<template>
    <div class="container mx-auto mb-5">
        <div class="shadow sm:rounded-md sm:overflow-hidden mb-5">
            <div class="grid grid-cols-3 gap-4 px-4 py-5 bg-white sm:p-6">
                <div class="form-control mt-0">
                    <label class="label">
                        <span class="label-text">数据库域名</span>
                    </label>
                    <input type="text" v-model="host" placeholder="127.0.0.1" class="input input-info input-bordered">
                    <label class="label">
                        <span class="label-text-alt">Please enter data</span>
                    </label>
                </div>
                <div class="form-control mt-0">
                    <label class="label">
                        <span class="label-text">数据库端口</span>
                    </label>
                    <input type="text" v-model="port" placeholder="3306" class="input input-info input-bordered">
                    <label class="label">
                        <span class="label-text-alt">Please enter port</span>
                    </label>
                </div>
                <div class="form-control mt-0">
                    <label class="label">
                        <span class="label-text">数据库名称</span>
                    </label>
                    <input type="text" v-model="dbname" placeholder="" class="input input-info input-bordered">
                    <label class="label">
                        <span class="label-text-alt">Please enter data</span>
                    </label>
                </div>
                <div class="form-control mt-0">
                    <label class="label">
                        <span class="label-text">数据库用户名</span>
                    </label>
                    <input type="text" v-model="username" placeholder="username" class="input input-info input-bordered">
                    <label class="label">
                        <span class="label-text-alt">Please enter data</span>
                    </label>
                </div>
                <div class="form-control mt-0">
                    <label class="label">
                        <span class="label-text">数据库密码</span>
                    </label>
                    <input type="password" v-model="password" placeholder="password" class="input input-info input-bordered">
                    <label class="label">
                        <span class="label-text-alt">Please enter data</span>
                    </label>
                </div>
            </div>

            <div class="px-4 py-3 bg-gray-50 text-right sm:px-6">
                <button class="btn btn-primary" @click="connHandle">连接数据库</button>
            </div>
        </div>
        <div class="shadow sm:rounded-md sm:overflow-hidden mb-5" v-if="tables.length > 0">
            <div class="px-4 py-5 bg-white sm:p-6">
            <ul class="accordion accordion-arrow" >
                <li class="accordion-item mb-3" v-for="(table,i) of tables" key="table.name">
                    <input v-model="tables[i].status" @click="toggleStatus(i)" :id="`item-793472${i}`" type="checkbox">
                    <label :for="`item-793472${i}`" class="text-xl font-medium accordion-title">
                        {{ table.name }}
                    </label>
                    <div class="accordion-body">
                        <div class="mockup-code" style="background-color: #282c34">
                            <pre v-highlightjs>
                                <code class="go">{{ table.res }}</code>
                            </pre>
                        </div>

                    </div>
                </li>
            </ul>
        </div>
        </div>
    </div>
</template>

<script>
import {defineComponent, reactive, toRefs, ref, onMounted} from 'vue'
import axios from "../plugins/axios";


export default defineComponent({
    setup(){
        const form = reactive({
            "host":"127.0.0.1",
            "port":3306,
            "dbname":"tinyurl",
            "username":"root",
            "password":"root"
        })

        const tables = ref([])

        const connHandle = ()=>{
            axios.post(`/conn`,form).then(({data})=>{
                if (data && data.length > 0) {
                    tables.value = data.map(v=>{
                        v.status = false
                        return v
                    })
                }

            })
            console.log(form)
        }

        const toggleStatus = i=>{
            console.log(i)
            tables.value[i].status = !tables.value[i].status
            axios.get(`/table/${tables.value[i].name}`).then(({data})=>{
                tables.value[i].res = data
            })
        }

        return {
            ...toRefs(form),
            connHandle,
            tables,
            toggleStatus,
        }
    }
})
</script>

<style scoped>

</style>