<template>
    <div class="console">
        <ul>
            <li v-for="log in logs">
                {{log}}
            </li>
        </ul>
    </div>
</template>

<script>
    export default {
        name: "Console",
        data() {
            return {
                logs : [ ],
                subbed : false
            };
        },

        beforeMount(){
            this.subscribeLogs();
        },

        methods: {
            subscribeLogs: function() {
                let self = this;
                if(!self.subbed) {
                    window.backend.MyClient.SubscribeLogs("log_message").then (result => {
                        wails.events.on("log_message", query => {
                            if (query) {
                                if (!self.logs.includes(query[0].Log))
                                    self.logs.push(query[0].Log);
                            }
                        });
                    });
                    self.subbed = true
                }
            }
        }


    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .console{
        grid-column: 1/-1;
        grid-row: 4/5;
        border: 1px solid white;
        text-align: left;
    }

    .button{
        display: grid;
        grid-template-columns: 1fr 150px 1fr;
        grid-template-rows: 1fr 65px 1fr;
    }

    .container{
        display: grid;
        grid-template-columns: 80px 1fr 80px;
        grid-template-rows: 80px 1fr 1fr 20px;
        height: 100%;
    }

    .value{
        grid-row : 1/2;
        grid-column: 2/3;
    }

    .item22{
        grid-row: 2/3;
        grid-column: 2/3;
    }

    ul {
        display: grid;
        list-style-type: none;
    }
    h1 {
    }
    a:hover {
        font-size: 1em;
        border-color: blue;
        background-color: blue;
        color: white;
        border: 3px solid white;
        border-radius: 10px;
        padding: 9px;
        cursor: pointer;
        transition: 500ms;
    }
    a {
        font-size: 1em;
        border-color: white;
        background-color: #121212;
        color: white;
        border: 3px solid white;
        border-radius: 10px;
        padding: 9px;
        cursor: pointer;
    }
</style>
