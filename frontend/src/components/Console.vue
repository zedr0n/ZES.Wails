<template>
    <div class="console container">
        <ul class="item1">
            <li v-for="log in logs">
                {{log}}
            </li>
        </ul>
        <div class="item2">
            {{branch}}
        </div>
    </div>
</template>

<script>
    export default {
        name: "Console",
        data() {
            return {
                logs : [ ],
                branch : " "
            };
        },

        beforeMount(){
            this.subscribeLogs();
            this.subscribeBranch();
        },

        methods: {
            subscribeLogs: function() {
                let self = this;

                window.backend.MyClient.SubscribeLogs("log_message").then (result => {
                    wails.events.on("log_message", query => {
                        if (query) {
                            if (!self.logs.includes(query[0].Log))
                                self.logs.push(query[0].Log);
                        }
                    });
                });
            },

            subscribeBranch : function () {
                let self = this;

                window.backend.MyClient.SubscribeBranch("active_branch").then (result => {
                    wails.events.on("active_branch", query => {
                        if (query) {
                            self.branch = query[0].ActiveBranch;
                        }
                    });
                });
            }
        },





    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .item1{
        grid-row: 1/2;
    }

    .item2{
        grid-row: 2/3;
        text-align: right;
    }

    .console{
        grid-column: 1/-1;
        grid-row: 4/5;
        border: 1px solid white;
        text-align: left;
    }

    .container{
        display: grid;
        grid-template-columns: 1fr;
        grid-template-rows: 1fr 20px;
        height: 100%;
    }

    ul {
        display: grid;
        list-style-type: none;
        overflow:hidden;
        height: 80%;
        overflow-y:scroll;

        margin-block-start: 1em;
        margin-block-end: 0em;
        margin-inline-start: 0px;
        margin-inline-end: 0px;
        padding-inline-start: 0px;
    }

    li {
        font-family: "Roboto";
        font-size: 12px;
    }

</style>
