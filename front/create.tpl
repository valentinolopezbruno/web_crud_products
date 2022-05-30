{{define "create"}}
{{template "header"}}


<div class="card">

    <div class="card-header">
        INGRESAR VALORES DEL NUEVO PRODUCTO
    </div>



    <div class="card-body">

        <form method="post" action="/insert">

        <div class="form-group">
            <label for=""> NOMBRE: </label>
            <input type="text"
            class="form-control" name="name" id="name" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

                <div class="form-group">
            <label for=""> DESCRIPCION: </label>
            <input type="text"
            class="form-control" name="description" id="description" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> URL IMAGEN: </label>
            <input type="text"
            class="form-control" name="img" id="img" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <button type="submit" class="btn btn-primary"> REGISTRAR PRODUCTO </button>

        <a name="" id="" class="btn btn-success" href="/home" role="button"> REGRESAR </a>

        





{{template "footer"}}

{{end}}