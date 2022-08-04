import React from 'react'

const SearchBox = () => {
    return (
        <>
            <div id="custom-search-input">
                <div className="input-group col-md-12">
                    <input
                        type="text"
                        className="  search-query form-control"
                        placeholder="Conversation"
                    />
                    <button className="btn btn-danger" type="button">
                        <span className=" glyphicon glyphicon-search" />
                    </button>
                </div>
            </div>
        </>
    )
}

export default SearchBox