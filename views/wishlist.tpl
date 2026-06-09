<div class="page-header">
    <h1>Travel Wishlist</h1>
    <p>Edit notes, update trip status, or remove destinations. Changes save without reloading the page.</p>
</div>

<div class="wishlist-wrapper">
    <table class="wishlist-table">
        <thead>
            <tr>
                <th>COUNTRY</th>
                <th>NOTE</th>
                <th>STATUS</th>
                <th>ACTIONS</th>
            </tr>
        </thead>
        <tbody id="wishlist-rows">
            {{range .WishlistItems}}
            <tr data-id="{{.ID}}">
                <td>{{.CountryName}}</td>
                <td><input type="text" class="note-input" value="{{.Note}}" placeholder="Add a note..."></td>
                <td>
                    <select class="status-select">
                        <option value="Planned" {{if eq .Status "Planned"}}selected{{end}}>Planned</option>
                        <option value="Visited" {{if eq .Status "Visited"}}selected{{end}}>Visited</option>
                    </select>
                </td>
                <td>
                    <button class="btn-save" onclick="saveRow(this)">Save</button>
                    <button class="btn-delete" onclick="deleteRow(this)">Delete</button>
                </td>
            </tr>
            {{else}}
            <tr><td colspan="4" style="text-align:center;color:#888;padding:2rem">No destinations yet. Add one from a country page!</td></tr>
            {{end}}
        </tbody>
    </table>
</div>

<script src="/static/js/wishlist.js"></script>
