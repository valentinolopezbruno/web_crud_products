{{define "edit"}}
{{template "header"}}


<div class="card">

    <div class="card-header">

        INGRESAR VALORES DEL PRODUCTO A MODIFICAR

    </div>

    <div class="card-body">

        <form method="post" action="/update">

        <div class="form-group">
            <label class="sr-only" for="inputName"> ID REGISTRO </label>
            <input type="text" value={{.Id}} class="form-control" name="id" id="id" placeholder="">
        </div>


        <div class="form-group">
            <label for=""> NOMBRE: </label>
            <input type="text"
            class="form-control" name="name" value={{.Name}} id="name" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

                <div class="form-group">
            <label for=""> DESCRIPCION: </label>
            <input type="text"
            class="form-control" name="description" value={{.Description}} id="description" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> URL DE IMAGEN: </label>
            <input type="text"
            class="form-control" name="img" value={{.Img}} id="img" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <button type="submit" class="btn btn-primary"  href="/delete?id={{.Id}}" > CONFIRMAR EDICIÓN  </button>

        <a name="" id="" class="btn btn-success" href="/home" role="button"> REGRESAR </a>

        
    </div>

    <div class="card-footer text-muted">
        Footer
    </div>

</div>






{{template "footer"}}
{{end}}