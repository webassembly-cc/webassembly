import React,{PropTypes} from 'react'
import ArticleList from '../components/article/list.jsx'
import { connect } from 'dva'


const Article = ({location,dispatch,article})=>{
    const {list}  = article

    console.info("router:-",list)
   
    const ArticleListProps = {
        list,
    }
    return(
        <div>
            <ArticleList {...ArticleListProps} />
        </div>
    )
}


Article.propTypes = {
    article : PropTypes.object
}

function mapStateToProps({ article }) {
  return { article };
}

export default connect(mapStateToProps)(Article)