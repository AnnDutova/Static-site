async function getData(url = '') {
    return await fetch(url, {
        method: 'GET',
    });
}

async function postData(url = '', data = {}) {
    return await fetch(url, {
        method: 'POST',
        body: JSON.stringify(data)
    });
}

async function callGo() {
    getData('http://localhost:8080/data')
        .then((data) => {
            data.json().then(data =>
                document.getElementById("response").innerHTML = data.field
            )
        });
}

async function createAccount(event) {
    event.preventDefault();

    current = new FormData(event.target);
    document.cookie = "username=" + current.get("username")
    postData('http://localhost:8080/createAccount', {
        username: current.get("username"),
        email: current.get("email"),
        password: current.get("password")
    })
        .then((data) => {
            data.json().then(data =>{
                if (data.error == undefined ){
                    window.location.replace("/user")
                }
                //document.getElementById("response").innerHTML = data.error
            })
        });
}

async function createSellerAccount(event) {
    event.preventDefault();
    current = new FormData(event.target);
    postData('http://localhost:8080/createSellerAccount', {
        username: current.get("username"),
        email: current.get("email"),
        password: current.get("password"), 
        salon_name: current.get("salon_name")
    })
        .then((data) => {
            data.json().then(data =>{
                if (data.error == undefined ){
                    window.location.replace("/seller")
                }
            })
        });
}


async function logIn(event) {
    event.preventDefault();
    current = new FormData(event.target);
    postData('http://localhost:8080/loginAccount', {
        username: current.get("username"),
        password: current.get("password")
    }).then((data) => {
        data.json().then( (data) => {
            if (data.error == undefined ){
                window.location.replace("/user")
            }else{
                document.getElementById("allert_message").innerHTML = data.error
                document.getElementById("allert_message").style.display = ""
            }
        })
    });
}

async function logInSeller(event) {
    event.preventDefault();
    current = new FormData(event.target);
    postData('http://localhost:8080/loginSellerAccount', {
        salon_name: current.get("salon_name"),
        password: current.get("password")
    }).then((data) => {
        data.json().then( (data) => {
            if (data.error == undefined ){
                window.location.replace("/seller")
            }
        })
    });
}

async function logInAdministrator(event){
    event.preventDefault();
    current = new FormData(event.target);
    console.log( current.get("username"))
    console.log( current.get("password"))
    postData('http://localhost:8080/loginAdministratorAccount', {
        username: current.get("username"),
        password: current.get("password")
    }).then((data) => {
        data.json().then( (data) => {
            if (data.error == undefined ){
                window.location.replace("/market")
            }
        })
    });
}


var stars = 0;

async function sendRewiew(event){
    event.preventDefault();
    postData('http://localhost:8080/sendRew', {
        rewiew: document.getElementById('userRewiew').value,
        author: document.getElementById('author_name').innerHTML,
        song: document.getElementById('song_title').innerHTML,
        salon: document.getElementById('salon_name').innerHTML,
        grade: stars
    }).then((data) => {
        document.getElementById('userRewiew').value = ""
        stars = 0;
        postData('http://localhost:8080/musicCard', {
            author: document.getElementById('author_name').innerHTML,
            song: document.getElementById('song_title').innerHTML,
            salon: document.getElementById('salon_name').innerHTML
        }).then((data) => {
            data.json().then((data) => {
                if (data.error == undefined ){
                    window.location.replace("/card")
                }
            })
        });
    });
}


async function sendPoint( stars_){
    stars = stars_;
}


async function createMusicCard(){
    event.preventDefault();
    console.log("Create")
    postData('http://localhost:8080/createMusicCard', {
        title: document.getElementById('title').value,
        duration: document.getElementById('duration').value,
        genre: document.getElementById('genre_inner').value,
        artist: document.getElementById('artists_inner').value,
        count: document.getElementById('count').value,
        price: document.getElementById('price').value,
        sale: document.getElementById('sale').value
    }).then((data) => {
        data.json().then((data) => {
            document.getElementById('title').value="",
            document.getElementById('duration').value="",
            document.getElementById('genre_inner').value="",
            document.getElementById('artists_inner').value="",
            document.getElementById('count').value ="",
            document.getElementById('price').value ="",
            document.getElementById('sale').value=""
        })
    });
}
function Answer (song, author, salon) {
    this.song = song;
    this.author = author;
    this.salon = salon;
}
async function createCollectionSeller(number){
    var selected = []
    var checked = []
    var song = ""
    var author =""
    for (let i = 0; i < number; i++) {
        console.log(document.getElementById(i).checked)
        if (document.getElementById(i).checked === true){
            checked.push(document.getElementById(i))
            var card = new Answer(
                document.getElementById('card_' + i).querySelector('h5').innerHTML,
                document.getElementById('card_' + i).querySelectorAll('p')[0].innerHTML,
                document.getElementById('card_' + i).querySelectorAll('p')[1].innerHTML
            )
            song += document.getElementById('card_' + i).querySelector('h5').innerHTML + ","
            author+=document.getElementById('card_' + i).querySelectorAll('p')[0].innerHTML+","
            salon += document.getElementById('card_' + i).querySelectorAll('p')[1].innerHTML+","
            selected.push(card)
        }
    }
    postData('http://localhost:8080/createCollection', {
        title: song,
        artist: author,
        collection_title: document.getElementById('collection_title').value,
        collection_price: document.getElementById('collection_price').value,
        collection_sale: document.getElementById('collection_sale').value
    }).then((data) => {
        data.json().then((data) => {
            console.log(data)
            for (let i = 0; i < number; i++) {
                document.getElementById(i).checked = false
            }
        })
    });
}

async function addMusicCardToBucket(event){
    event.preventDefault();
    postData('http://localhost:8080/addMusicCardToBucketFromCollection', {
        author: document.getElementById('author_name').textContent,
        song: document.getElementById('song_title').textContent,
        salon: document.getElementById('salon_name').textContent
    })
}

async function addMusicCardToBucketFromCollection(song_, author_, salon_){
    postData('http://localhost:8080/addMusicCardToBucketFromCollection', {
        author: author_,
        song: song_,
        salon: salon_
    })
}

async function addAllCollectionToBucket(title_, salon_){
    console.log(title_,  salon_)
    postData('http://localhost:8080/addAllCollectionToBucket', {
        title: title_,
        salon: salon_
    }).then((data) => {
        data.json().then((data) => {
            console.log(data)
        })
    });
}

async function openCompositionCardFromSeller(song_, author_, salon_){
    postData('http://localhost:8080/musicCardSeller', {
        author: author_,
        song: song_,
        salon: salon_
    }).then((data) => {
        data.json().then((data) => {
            if (data.error == undefined ){
                window.location.replace("/cardSeller")
            }
        })
    });
}

async function DeliteElementFromBucket(song_, author_, salon_){
    postData('http://localhost:8080/deliteFromBucket', {
        author: author_,
        song: song_,
        salon: salon_
    })
}


async function openMusicCard(song_, author_, salon_){
    postData('http://localhost:8080/musicCard', {
        author: author_,
        song: song_,
        salon: salon_
    }).then((data) => {
        data.json().then((data) => {
            console.log("openMusicCard")
            if (data.error == undefined ){
                window.location.replace("/card")
            }
        })
    });
}
async function BlockUserAdmin(user){
    postData('http://localhost:8080/blockUser', {
        username: user
    })
}
async function DeliteRewiewAdmin(user, grade, text_){
    postData('http://localhost:8080/deliteRewiew', {
            username: user,
            grade: grade, 
            text: text_
        })
}

async function openCollectionCard(title_, value_, salon_){
        postData('http://localhost:8080/collectionCard', {
            title: title_,
            value: value_, 
            salon: salon_
        }).then((data) => {
            data.json().then((data) => {
                if (data.error == undefined ){
                    window.location.replace("/card")
                }
            })
        });
}

async function openCompositionCard(isAdmin, song_, author_, salon_){
    console.log("openCompositionCard")
    if (isAdmin == null){
        postData('http://localhost:8080/musicCard', {
            author: author_,
            song: song_,
            salon: salon_
        }).then((data) => {
            data.json().then((data) => {
                if (data.error == undefined ){
                    window.location.replace("/card")
                }
            })
        });
    }else{
        console.log(song_)
        postData('http://localhost:8080/openCompositionPageFromAdministrator', {
            author: author_,
            song: song_,
            salon: salon_
        }).then((data) => {
            data.json().then((data) => {
                if (data.error == undefined ){
                    window.location.replace("/cardFromAdmin")
                }
            })
        });
    }
}

async function openCardFromMarket(isAdmin, song_, author_, salon_){
    console.log("openCardFromMarket")
    if (isAdmin == null){
        postData('http://localhost:8080/musicCard', {
            author: author_,
            song: song_,
            salon: salon_
        }).then((data) => {
            data.json().then((data) => {
                if (data.error == undefined ){
                    window.location.replace("/card_")
                }
            })
        });
    }else{
        console.log(song_)
        postData('http://localhost:8080/openCompositionPageFromAdministrator', {
            author: author_,
            song: song_,
            salon: salon_
        }).then((data) => {
            data.json().then((data) => {
                if (data.error == undefined ){
                    window.location.replace("/cardFromAdmin")
                }
            })
        });
    }
}

async function BuyAllBucket(data){
    getData('http://localhost:8080/buyFromBucketCard')
    .then((data) => {
        data.json().then(data =>{
            console.log(data)
            if (data.error != undefined ){
                document.getElementById("error_in_transaction").textContent = data.error
                document.getElementById("error_in_transaction").style.display = ""
            }
        })
    });
}

async function AddNewPreference(number){
    var preference = ""
    for (let i = 0; i < number; i++) {
        console.log(document.getElementById("checkbox_value_"+i).checked)
        if (document.getElementById("checkbox_value_"+i).checked === true){
            preference += (i + 1) + ","
        }
    }
    postData('http://localhost:8080/AddPreference', {
        preferences: preference
    }).then((data) => {
        data.json().then( (data) => {
            console.log("Add preference")
            while (document.getElementById("preference_type_el") != null){
                document.getElementById("preference_type_el").remove()
            }
            document.getElementById("add_pref_form").remove()
            document.getElementById("add_new_pref").remove()
        })
    });
}

async function openBucketCard(song_, author_, salon_){
    postData('http://localhost:8080/bucketCard', {
        author: author_,
        song: song_,
        salon: salon_
    }).then((data) => {
        data.json().then( (data) => {
            if (data.error == undefined ){
                window.location.replace("/card")
            }
        })
    });
}

async function replenish(event){
    event.preventDefault();
    current = new FormData(event.target);
    postData('http://localhost:8080/replenishWallet', {
        count: current.get("money")
    }).then((data) => {
        data.json().then( (data) => {
            console.log(data)
            if (data.error == undefined ){
                window.location.replace("/user")
            }
        })
    });
}

async function createSale(event){
    event.preventDefault();
    current = new FormData(event.target);
    postData('http://localhost:8080/addSaleInDB', {
        count: current.get("count")
    }).then((data) => {
        data.json().then( (data) => {
            if (data.error == undefined ){
                console.log(data)
                if (data.id_user == 9){
                    window.location.replace("/cardFromAdmin")
                }else{
                    window.location.replace("/cardSeller")
                }
            }
        })
    });
}

async function eraseCookie(name) {
    document.cookie = name + '=; Max-Age=-99999999;';
}
