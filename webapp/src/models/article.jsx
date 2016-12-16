import {query} from '../services/article'
import {parse} from 'qs'


export default{
    namespace:'article',

    state:{
        list :[]
    },
     subscriptions: {
        setup({ dispatch, history }) {
        history.listen(location => {
            if (location.pathname === '/') {
            dispatch({
                type: 'query',
                payload: location.query,
            });
            }
        });
        },
    },
    effects:{
        *query({payload},{call,put}){
            const {data}  =yield call(query,parse(payload))

 console.info(data)

            if(data){
                yield put({
                    type:'querySuccess',
                    payload:{
                        list:data.data
                    }
                })
            }
        }
    },

    reducers:{
        querySuccess(state,action){
            return {...state,...action.payload}
        }
    }

}